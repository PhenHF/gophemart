package database

import (
	"context"

	commonTypes "github.com/PhenHF/gophemart/internal/common"
)

func (db *DataBase) CreateNewUser(ctx context.Context, user commonTypes.User) (uint, error) {
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}

	query := `INSERT INTO user (login, password) VALUES ($1, $2)`

	_, err = tx.ExecContext(ctx, query, user.Login, user.Password)
	if err != nil {
		return 0, err
	}

	userID := db.GetUserID(ctx, user)
	
	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (db *DataBase) GetUserID(ctx context.Context, user commonTypes.User) uint {
	query := `SELECT id FROM user WHERE login=$1 AND passowrd=$2`
	row := db.QueryRowContext(ctx, query, user.Login, user.Password)
	var userID uint

	err := row.Scan(&userID)
	if err != nil {
		return 0
	}

	return userID


}