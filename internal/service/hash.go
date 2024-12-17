package service

import (
	"crypto/sha256"
	"fmt"
	"sync"

	commonTypes "github.com/PhenHF/gophemart/internal/common"
)

func HashSumUserCreds(user *commonTypes.User) {
	var wg sync.WaitGroup
	wg.Add(2)

	bLogin := []byte(user.Login + "VERY_SECRET_TOKEN")
	bPassword := []byte(user.Password + "VERY_SECRET_TOKEN")

	go func() {
		h := sha256.New()
		h.Write(bLogin)
		user.Login = fmt.Sprintf("%x", h.Sum(nil))
		wg.Done()
	}()

	go func() {
		h := sha256.New()
		h.Write(bPassword)
		user.Password = fmt.Sprintf("%x", h.Sum(nil))
		wg.Done()
	}()
		
	wg.Wait()
}
