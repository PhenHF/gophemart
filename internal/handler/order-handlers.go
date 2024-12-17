package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/PhenHF/gophemart/internal/common"
	"github.com/PhenHF/gophemart/internal/database"
	"github.com/PhenHF/gophemart/internal/service"
	auth "github.com/PhenHF/gophemart/pkg/jwtauth"
)

func UploadUserOrder(storage common.Storager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {		
		userID := auth.CheckAuth(r)

		if userID == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if len(body) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		
		orderInt := service.ConvertBodyToInt(body)
		err = service.ValidateOrder(r.Context(), userID, orderInt, storage)
		if err != nil {
			switch {
			case errors.As(err, &service.InvalidFromatOrderError):
				fmt.Println(err)
				w.WriteHeader(http.StatusUnprocessableEntity)
				return
			case errors.As(err, &database.OrderAlreadyExistsError):
				fmt.Println(err)
				w.WriteHeader(http.StatusOK)
				return
			
			case errors.As(err, &database.OrderAlreadyExistsForAnotherUserError):
				fmt.Println(err)
				w.WriteHeader(http.StatusConflict)
				return

			default:
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		// #TODO implement send request on /api/orders/{number}

		order := &common.Order{
			Number: uint(orderInt),
			Status: "OK",
			UserID: userID,
		}

		err = storage.InsertOrder(r.Context(), *order)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)
	})
}

func GetUserOrders(storage common.Storager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		orders := make([]common.Order, 0)

		userID := auth.CheckAuth(r)

		err := storage.SelectAllUserOrders(r.Context(), &orders, userID)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if len(orders) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		response, err := json.Marshal(orders)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)

	})
}