package server

import (
	"net/http"
)

func RunServer() {
	rt := buildRt()
	if err := http.ListenAndServe(`:8000`, rt); err != nil {
		panic(err)
	} 
}
