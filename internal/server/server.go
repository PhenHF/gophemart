package server

import (
	"net/http"

	"github.com/PhenHF/gophemart/internal/common"
)

func RunServer(storage common.Storager) {
	rt := buildRt(storage)
	if err := http.ListenAndServe(`:8000`, rt); err != nil {
		panic(err)
	}
}
