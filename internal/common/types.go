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
	UploadedAt string `json:"uploaded_at"`
	UserID uint
}

type Storager interface {
	InsertNewUser(ctx context.Context, user User) (uint, error)
	SelectUserID(ctx context.Context, user User) uint
	CheckOrderInDB(ctx context.Context, order int, userID uint) error
	InsertOrder(ctx context.Context, order Order) error 
}
