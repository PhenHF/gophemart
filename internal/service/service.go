package service

import "strconv"

func convertBodyToInt(order []byte) int {
	orderInt, err := strconv.Atoi(string(order)) 
	if err != nil {
		return 0
	}
	return orderInt
}