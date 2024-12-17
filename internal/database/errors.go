package database

import "fmt"

type OrderAlreadyExistsForAnotherUser struct {
	OrderNum int
	Err error
}

func (oe *OrderAlreadyExistsForAnotherUser) Error() string {
	return fmt.Sprintf("order number %v was uploaded by another user", oe.OrderNum)
}

func NewOrderAlreadyExistsForAnotherUser(order int, err error) error {
	return &OrderAlreadyExistsForAnotherUser{
		OrderNum: order,
		Err: err,
	}
}

type OrderAlreadyExists struct {
	OrderNum int
	Err error
}

func (oe *OrderAlreadyExists) Error() string {
	return fmt.Sprintf("order number %v was uploaded", oe.OrderNum)
}

func NewOrderAlreadyExists(order int, err error) error {
	return &OrderAlreadyExists{
		OrderNum: order,
		Err: err,
	}
}

var OrderAlreadyExistsError *OrderAlreadyExists
var OrderAlreadyExistsForAnotherUserError *OrderAlreadyExistsForAnotherUser