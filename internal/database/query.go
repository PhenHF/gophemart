package database

import (
	"context"

	"github.com/PhenHF/gophemart/internal/common"
)

func (db *DataBase) InsertNewUser(ctx context.Context, user common.User) (uint, error) {
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}

	query := `INSERT INTO user (login, password) VALUES ($1, $2)`

	_, err = tx.ExecContext(ctx, query, user.Login, user.Password)
	if err != nil {
		return 0, err
	}

	userID := db.SelectUserID(ctx, user)

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (db *DataBase) SelectUserID(ctx context.Context, user common.User) uint {
	query := `SELECT id FROM user WHERE login=$1 AND passowrd=$2`
	row := db.QueryRowContext(ctx, query, user.Login, user.Password)
	var userID uint

	err := row.Scan(&userID)
	if err != nil {
		return 0
	}

	return userID

}

func (db *DataBase) IsertOrder(ctx context.Context, ch chan bool) (bool, error) {
	// query := `INSERT INTO order (number, status, uploaded_at, user_id, accrual) 
	// 		VALUES($1, $2, $3, $4, $5)`
	
	query := `SELECT user_id FROM order WHERE nubmer=$1`

	row := db.QueryRowContext(ctx, query)

	if ok := <- ch; !ok {
		
	}

	return true, nil
}