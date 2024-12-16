package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	auth "github.com/PhenHF/gophemart/pkg/jwtauth"
)

func UserRegistration(w http.ResponseWriter, r *http.Request) {
	uc := userCreds{}

	if err := json.NewDecoder(r.Body).Decode(&uc); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if uc.Login == "" || uc.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// #TODO implement adding a new user to the database
	
	token, err := auth.CreateNewJWTToken(1)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{Name: "user_id", Value: token}
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	uc := userCreds{}
	if err := json.NewDecoder(r.Body).Decode(&uc); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// #TODO implement JWT generation and cookie setting

	w.WriteHeader(http.StatusOK)
}
