package server

import (
	"net/http"

	commonTypes "github.com/PhenHF/gophemart/internal/common"
)

func RunServer(storage commonTypes.Storager) {
	rt := buildRt(storage)
	if err := http.ListenAndServe(`:8000`, rt); err != nil {
		panic(err)
	}
}
