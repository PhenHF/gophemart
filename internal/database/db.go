package database

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
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
	
	createTables(storage)

	return &DataBase{storage}
}

func createTables(storage *sql.DB) {
	tx, err := storage.Begin()
	if err != nil {
		panic(err)
	}
	
	query := `CREATE TABLE IF NOT EXISTS  users (
		"id" SERIAL PRIMARY KEY,
		"login" text UNIQUE NOT NULL,
		"password" text NOT NULL
	)`

	_, err = tx.Exec(query)
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	query = `CREATE TABLE IF NOT EXISTS  orders (
		"id" SERIAL PRIMARY KEY,
		"number" bigint UNIQUE NOT NULL,
		"status" varchar(20) NOT NULL,
		"accrual" int, 
		"uploaded_at" text NOT NULL,
		"user_id" bigint REFERENCES users
	)`

	_, err = tx.Exec(query)
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	query = `CREATE TABLE IF NOT EXISTS withdrawal (
		"id" SERIAL PRIMARY KEY,
		"number" bigint UNIQUE NOT NULL,
		"sum" bigint,
		"processed_at" text,
		"user_id" bigint REFERENCES users
	)`

	_, err = tx.Exec(query)
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	query = `CREATE TABLE IF NOT EXISTS balance (
		"id" SERIAL PRIMARY KEY,
		"sum" bigint,
		"user_id" bigint REFERENCES users
	)`

	_, err = tx.Exec(query)
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	tx.Commit()
}
