package database

import (
	"context"
	"database/sql"
	"time"
)

func NewDataBaseConnection() *DataBase {
	// TODO implement reading creds for db from environment variable
	storage, err := sql.Open("pgx", "host=localhost user=postgres password=1111 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err = storage.PingContext(ctx); err != nil {
		panic(err)
	}

	return &DataBase{storage}
}
