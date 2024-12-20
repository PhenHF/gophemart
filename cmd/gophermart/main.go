package main

import (
	"github.com/PhenHF/gophemart/internal/database"
	"github.com/PhenHF/gophemart/internal/server"
	"github.com/PhenHF/gophemart/internal/service"
)

func main() {
	storage := database.NewDataBaseConnection()
	go service.WorkersPool(storage)
	server.RunServer(storage)
}
