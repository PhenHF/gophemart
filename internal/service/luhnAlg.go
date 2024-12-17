package service

import "fmt"

 func LuhnValid(order int) error {
	if b := (order % 10 + checksum(order/10)) % 10 == 0; !b {
		return NewInvalidFormatOrder(order, nil)
	}

 	return nil
 }
 
 func checksum(order int) int {
 	var luhn int
 
 	for i := 0; order > 0; i++ {
 		cur := order % 10
 
 		if i % 2 == 0 { // even
 			cur = cur * 2
 			if cur > 9 {
 				cur = cur % 10 + cur / 10
 			}
 		}
 
 		luhn += cur
 		order = order / 10
 	}
 	return luhn % 10
}

type InvalidFormatOrder struct {
	Order int
	Err error
}

func (io *InvalidFormatOrder) Error() string {
	return fmt.Sprintf("invalid format order %v", io.Order)
}

func NewInvalidFormatOrder(order int, err error) error {
	return &InvalidFormatOrder{
		Order: order,
		Err: err,
	}
}

var InvalidFromatOrderError *InvalidFormatOrder