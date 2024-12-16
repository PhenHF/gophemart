package main

import (
	"github.com/PhenHF/gophemart/internal/database"
	"github.com/PhenHF/gophemart/internal/server"
)

func main() {
	storage := database.NewDataBaseConnection()
	server.RunServer(storage)
}
