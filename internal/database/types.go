package database

import "database/sql"

type DataBase struct {
	*sql.DB
}