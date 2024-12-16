package common

import "context"

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UploadOrder struct {
	Number uint `json:"number"`
}

type Storager interface {
	InsertNewUser(ctx context.Context, user User) (uint, error)
	SelectUserID(ctx context.Context, user User) uint
}
