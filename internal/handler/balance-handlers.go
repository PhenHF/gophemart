package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PhenHF/gophemart/internal/common"
	"github.com/PhenHF/gophemart/internal/service"
	auth "github.com/PhenHF/gophemart/pkg/jwtauth"
)

func GetUserBalance(storage common.Storager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := auth.CheckAuth(r)
		if userID == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		balance := &common.Balance{}

		err := storage.SelectCurrentBalance(r.Context(), userID, *balance)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response, err := json.Marshal(balance)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	})
}

func WriteOffPointsForPayMents(storage common.Storager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := auth.CheckAuth(r)
		if userID == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var WriteOffReques struct {
			Order uint `json:"order"`
			Sum uint `json:"sum"`
		}

		if err := json.NewDecoder(r.Body).Decode(&WriteOffReques); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		} 
		
		err := service.LuhnValid(int(WriteOffReques.Order))
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		err = storage.UpdatePointsForAnOrders(r.Context(), userID, WriteOffReques.Order, WriteOffReques.Sum)
		if err != nil {
			// #TODO implemen switch case for errs
			return
		}
		
		w.WriteHeader(http.StatusOK)
	})
}

func GetInfoAboutWithdrawal(storage common.Storager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := auth.CheckAuth(r)
		if userID == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		withdrawals := make([]common.Withdrawal, 0)

		if err := storage.SelectAllUsersWithdrawals(r.Context(), userID, &withdrawals); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if len(withdrawals) == 0 {
			w.WriteHeader(http.StatusNoContent)
		}


		res, err := json.Marshal(withdrawals)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	})
}