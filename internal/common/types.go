package common

import (
	"context"
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Order struct {
	Number uint `json:"number"`
	Status string `json:"status"`
	Accrual int `json:"accrual,omitempty"`
	UploadedAt string `json:"uploaded_at,omitempty"`
	UserID uint `json:"-"`
}

type Balance struct {
	Current uint `json:"current"`
	Withdrawn uint `json:"withdrawn"`
}

type Withdrawal struct {
	Number uint `json:"order"`
	Sum uint `json:"sum"`
	ProcessedAt string `json:"processed_at"`
}

type Storager interface {
	InsertNewUser(ctx context.Context, user User) (uint, error)
	SelectUserID(ctx context.Context, user User) uint
	CheckOrderInDB(ctx context.Context, order int, userID uint) error
	InsertOrder(ctx context.Context, order Order) error 
	SelectAllUserOrders(ctx context.Context, orders *[]Order, userID uint) error
	SelectCurrentBalance(ctx context.Context, userID uint, balance *Balance) error
	UpdatePointsForAnOrders(ctx context.Context, userID, order, sum uint) error
	SelectAllUsersWithdrawals(ctx context.Context, userID uint, withdrawals *[]Withdrawal) error
	UpdateOrder(ctx context.Context, order Order) error
	UpdateBalance(ctx context.Context, userID uint, sum uint) error
} 
