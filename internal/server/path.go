package server

import (
	commonTypes "github.com/PhenHF/gophemart/internal/common"
	"github.com/PhenHF/gophemart/internal/handler"
	"github.com/go-chi/chi/v5"
)

func buildRt(storage commonTypes.Storager) *chi.Mux{
	rt := chi.NewRouter()

	rt.Post(`/api/user/register`, handler.UserRegistration(storage))
	rt.Post(`/api/user/login`, handler.UserLogin(storage))
	return rt
}