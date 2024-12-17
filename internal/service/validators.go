package service

import (
	"context"

	"github.com/PhenHF/gophemart/internal/common"
)

func ValidateOrder(ctx context.Context, userID uint, order int, storage common.Storager) error {
	errCh := make(chan error)

	go func(){
		errCh <- storage.CheckOrderInDB(ctx, order, userID)
	}()

	go func() {
		errCh <- LuhnValid(order)
	}()

	for i := 0; i < 2; i ++ {
		errFromCh := <- errCh
		if errFromCh != nil {
			return errFromCh
		}
	}
	
	return nil
}