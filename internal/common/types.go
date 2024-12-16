package common

import "context"

type User struct {
	Login string `json:"login"`
	Password string `json:"password"`
}

type Storager interface {
	CreateNewUser(ctx context.Context, user User) (uint, error)
	GetUserID(ctx context.Context, user User) uint
}