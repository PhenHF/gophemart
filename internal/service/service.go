package service

import (
	"fmt"
	"strconv"
)

func ConvertBodyToInt(order []byte) int {
	orderInt, err := strconv.ParseUint(string(order), 10, 64)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return int(orderInt)
}