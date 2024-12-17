package jwtauth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateNewJWTToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add((time.Hour * 3))),
		},

		UserID: userID,
	})

	// #TODO implement reading SECTET_TOKEN from an environment variable
	tokenStr, err := token.SignedString([]byte("VERY_SECRET_TOKEN"))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func PasreJWTToken(tokenStr string) (Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		// #TODO implement reading SECTET_TOKEN from an environment variable
		return []byte("VERY_SECRET_TOKEN"), nil
	})

	if err != nil {
		return *claims, err
	}

	if !token.Valid {
		return *claims, newInvalidTokenError(nil)
	}

	return *claims, nil

}

func Authtorization(userID uint, w http.ResponseWriter) error {
	token, err := CreateNewJWTToken(userID)
	if err != nil {
		// #TODO add err logging
		return err
	}

	cookie := &http.Cookie{Name: "Authtorization", Value: token}
	http.SetCookie(w, cookie)
	
	return nil
}  

func CheckAuth(r *http.Request) uint {
	token, err := r.Cookie("Authtorization")
	if err != nil {
		return 0
	}

	c, err := PasreJWTToken(token.Value)
	if err != nil {
		return 0
	}

	return c.UserID
}