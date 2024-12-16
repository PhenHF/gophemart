package service

func Valid(ch chan bool, order []byte) {
	convertBodyToInt(order)

	ch <- true
}