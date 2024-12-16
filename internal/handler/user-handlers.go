package handler

import (
	"encoding/json"
	"net/http"

	"github.com/PhenHF/gophemart/internal/common"
	"github.com/PhenHF/gophemart/internal/service"
	auth "github.com/PhenHF/gophemart/pkg/jwtauth"
)

func UserRegistration(storage common.Storager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uc := common.User{}
		if err := json.NewDecoder(r.Body).Decode(&uc); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if uc.Login == "" || uc.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		service.HashSumUserCreds(&uc)

		userID, err := storage.InsertNewUser(r.Context(), uc)
		if err != nil {
			// #TODO add err logging
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := auth.Authtorization(userID, w); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}
func UserLogin(storage common.Storager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uc := common.User{}
		if err := json.NewDecoder(r.Body).Decode(&uc); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		service.HashSumUserCreds(&uc)

		userID := storage.SelectUserID(r.Context(), uc)
		if userID == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		
		if err := auth.Authtorization(userID, w); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}
