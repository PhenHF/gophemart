package server

import (
	"github.com/PhenHF/gophemart/internal/common"
	"github.com/PhenHF/gophemart/internal/handler"
	"github.com/go-chi/chi/v5"
)

func buildRt(storage common.Storager) *chi.Mux {
	rt := chi.NewRouter()

	rt.Post(`/api/user/register`, handler.UserRegistration(storage))
	rt.Post(`/api/user/login`, handler.UserLogin(storage))
	rt.Post(`/api/user/orders`, nil)
	return rt
}
