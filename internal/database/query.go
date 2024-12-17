package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/PhenHF/gophemart/internal/common"
)

func (db *DataBase) InsertNewUser(ctx context.Context, user common.User) (uint, error) {
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}

	query := `INSERT INTO users (login, password) VALUES ($1, $2)`

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
	query := `SELECT id FROM users WHERE login=$1 AND password=$2`
	row := db.QueryRowContext(ctx, query, user.Login, user.Password)
	var userID uint

	err := row.Scan(&userID)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return userID

}

func (db *DataBase) CheckOrderInDB(ctx context.Context, order int, userID uint) error {
	query := `SELECT user_id FROM orders WHERE number=$1`

	var userIDFromOrder uint
	err := db.QueryRowContext(ctx, query, order).Scan(&userIDFromOrder)
	switch {
	case err == sql.ErrNoRows:
		return nil
		
	case err != nil:
		return err
	}

	if userID != userIDFromOrder {
		return NewOrderAlreadyExistsForAnotherUser(order, nil)
	}
	
	return NewOrderAlreadyExists(order, nil)
}

func (db *DataBase) InsertOrder(ctx context.Context, order common.Order) error {

	query := `INSERT INTO orders (number, status, user_id, accrual, uploaded_at) 
			VALUES($1, $2, $3, $4, $5)`

	_, err := db.ExecContext(
		ctx, query, 
		order.Number, 
		order.Status, 
		order.UserID, 
		order.Accrual, 
		time.Now().Format(time.RFC3339),
	)

	if err != nil {
		return err
	}
			
	return nil
}

