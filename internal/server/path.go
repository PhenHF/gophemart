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
	rt.Post(`/api/user/orders`, handler.UploadUserOrder(storage))
	rt.Get(`/api/user/orders`, handler.GetUserOrders(storage))
	rt.Post(`/api/user/balance/withdraw`, handler.WriteOffPointsForPayMents(storage))
	rt.Get(`/api/user/withdrawals`, handler.GetInfoAboutWithdrawal(storage))
	return rt
}
