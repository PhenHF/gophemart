package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/PhenHF/gophemart/internal/common"
)


func NewQueueNewOrdersCh() chan common.Order {
	return make(chan common.Order, 1)
}

var QueueNewOrderCh = NewQueueNewOrdersCh()

func worker(storage common.Storager, queueOrders chan common.Order) {
	for {
		o := <- queueOrders
		
		endpoint := "http://localhost:8080/api/orders/" + string(o.Number)
		
		client := &http.Client{}
		
		request, err := http.NewRequest(http.MethodGet, endpoint, nil)
		if err != nil {
			queueOrders <- o
			break
		}
		
		res, err := client.Do(request)
		if err != nil {
			queueOrders <- o
			continue
		}
		
		err = json.NewDecoder(res.Body).Decode(&o)
		if err != nil {
			queueOrders <- o
			continue
		}
		

		switch o.Status {
		case "PROCESSED":
			err = storage.UpdateOrder(context.Background(), o)
			if err != nil {
				fmt.Println(err)
			}
			err = storage.UpdateBalance(context.Background(), o.UserID, uint(o.Accrual))
			if err != nil {
				fmt.Println(err)
			}
		}

		if o.Status != "PROCESSED" && o.Status != "INVALID" {
			queueOrders <- o
		}
		
		fmt.Printf("Status: %v accrual: %v \n", o.Status, o.Accrual)
		err = storage.UpdateOrder(context.Background(), o)
		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(5 * time.Second)
	}
}


func WorkersPool(storage common.Storager) {
	queueOrders := make(chan common.Order, 10)
	go func() {
		for {
			order := <- QueueNewOrderCh
			queueOrders <- order
		}
	}()

	for w := 1; w < 3; w++ {
		go worker(storage, queueOrders)
	}
}