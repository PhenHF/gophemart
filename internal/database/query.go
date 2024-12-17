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
	defer tx.Rollback()
	if err != nil {
		return 0, err
	}

	query := `INSERT INTO users (login, password) VALUES ($1, $2)`

	_, err = tx.ExecContext(ctx, query, user.Login, user.Password)
	if err != nil {
		return 0, err
	}

	query = `SELECT id FROM users WHERE login=$1 AND password=$2`
	var userID uint
	err = tx.QueryRowContext(ctx, query, user.Login, user.Password).Scan(&userID)
	if err != nil && err != sql.ErrNoRows{
		return 0, nil
	}

	query = `INSERT INTO balance (sum, user_id) VALUES ($1, $2)`
	_, err = tx.ExecContext(ctx, query, 0, userID)
	if err != nil {
		return 0, err
	}

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

func (db *DataBase) SelectAllUserOrders(ctx context.Context, orders *[]common.Order, userID uint) error {
	query := `SELECT number, status, accrual, uploaded_at FROM orders WHERE user_id=$1`

	rows, err := db.QueryContext(ctx, query, userID)
	if err != nil && err != sql.ErrNoRows{
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var o common.Order
		err = rows.Scan(&o.Number, &o.Status, &o.Accrual, &o.UploadedAt)
		if err != nil {
			continue
		}

		*orders = append(*orders, o)
	}

	err = rows.Err()
	if err != nil {
		return err
	}

	return nil
}

func (db *DataBase) UpdateBalance(ctx context.Context, userID uint, sum uint) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `SELECT sum FROM balance WHERE user_id=$1`
	var userBalance uint
	err = tx.QueryRowContext(ctx, query, userID).Scan(&userBalance)
	if err != nil {
		return err
	}

	query = `UPDATE balance SET sum=$1 WHERE user_id=$2`
	
	userBalance += sum
	_, err = tx.ExecContext(ctx, query, userBalance, userID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (db *DataBase) SelectCurrentBalance(ctx context.Context, userID uint, balance common.Balance) error {
	query := `SELECT balance.sum, SUM(withdrawal.sum) FROM balance
			JOIN withdrawal ON balance.user_id = withdrawal.user_id
			GROUP BY balance.sum, balance.user_id
			HAVING balance.user_id=$1`
	
	err := db.QueryRowContext(ctx, query, userID).Scan(&balance.Current, balance.Withdrawn)
	if err != nil {
		return nil
	}

	return nil
} 

func (db *DataBase) UpdatePointsForAnOrders(ctx context.Context, userID, order, sum uint) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `SELECT sum FROM balance WHERE user_id=$1`
	var userBalance uint
	err = tx.QueryRowContext(ctx, query, userID).Scan(&userBalance)
	if err != nil {
		return err
	}

	if userBalance < sum {
		return NewSumGreaterBalance(nil)
	}

	query = `UPDATE balance SET sum=$1 WHERE user_id=$2`
	_, err = tx.ExecContext(ctx, query, userBalance - sum, userID)
	if err != nil {
		return err
	}

	query = `INSERT INTO withdrawal (number, sum, processed_at, user_id)`
	_, err = tx.ExecContext(ctx, query, order, sum, time.Now().Format(time.RFC3339), userID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (db *DataBase) SelectAllUsersWithdrawals(ctx context.Context, userID uint, withdrawals *[]common.Withdrawal) error {
	query := `SELECT number, sum, processed_at FROM withdrawal WHERE user_id=$1`
	
	rows, err := db.QueryContext(ctx, query, userID)
	if err != nil && err != sql.ErrNoRows{
		return err
	}
	defer rows.Close()
	
	for rows.Next() {
		var w common.Withdrawal
		err = rows.Scan(&w.Number, &w.Sum, &w.ProcessedAt)
		if err != nil {
			continue
		}
		
		*withdrawals = append(*withdrawals, w)
	}
	
	err = rows.Err()
	if err != nil {
		return err
	}
	
	return nil
}