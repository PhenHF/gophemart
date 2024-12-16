package handler

import (
	"encoding/json"
	"net/http"

	commonTypes "github.com/PhenHF/gophemart/internal/common"
	auth "github.com/PhenHF/gophemart/pkg/jwtauth"
)

func UserRegistration(storage commonTypes.Storager) http.HandlerFunc {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		uc := commonTypes.User{}
		if err := json.NewDecoder(r.Body).Decode(&uc); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if uc.Login == "" || uc.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		userID, err := storage.CreateNewUser(r.Context(), uc)
		if err != nil {
			// #TODO add err logging
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		token, err := auth.CreateNewJWTToken(userID)
		if err != nil {
			// #TODO add err logging
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		cookie := &http.Cookie{Name: "user_id", Value: token}
		http.SetCookie(w, cookie)
		w.WriteHeader(http.StatusOK)
	})		
}
func UserLogin(storage commonTypes.Storager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uc := commonTypes.User{}
		if err := json.NewDecoder(r.Body).Decode(&uc); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		userID := storage.GetUserID(r.Context(), uc)
		if userID == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err := auth.CreateNewJWTToken(userID)
		if err != nil {
			// #TODO add err logging
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		cookie := &http.Cookie{Name: "user_id", Value: token}
		http.SetCookie(w, cookie)
		w.WriteHeader(http.StatusOK)
	})
}

