package handler

import (
	"net/http"

	"github.com/PhenHF/gophemart/internal/common"
	"github.com/PhenHF/gophemart/internal/service"
	auth "github.com/PhenHF/gophemart/pkg/jwtauth"
)

func UploadUserOrder(storage common.Storager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var order []byte

		if userID := auth.CheckAuth(r); userID == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		res, err := r.Body.Read(order)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if res == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		lunaTest := make(chan bool)

		go service.Valid(lunaTest, order)


	})
}