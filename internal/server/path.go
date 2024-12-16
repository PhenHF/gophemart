package server

import (
	"github.com/PhenHF/gophemart/internal/handler"
	"github.com/go-chi/chi/v5"
)

func buildRt() *chi.Mux{
	rt := chi.NewRouter()

	rt.Post(`/api/user/register`, handler.UserRegistration)
	rt.Post(`/api/user/login`, handler.UserLogin)
	return rt
}