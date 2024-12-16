package jwtauth

import (
	"fmt"
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

func PasreJWTToken(tokenStr string) (Claims, error){
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		// #TODO implement reading SECTET_TOKEN from an environment variable
		return []byte("VERY_SECRET_TOKEN"), nil
	})

	if err != nil {
		return Claims{}, err
	}

	if !token.Valid {
		// #TODO implement errorType for invalid token
		return Claims{}, fmt.Errorf("token is not valid")
	}

	return *claims, nil

}