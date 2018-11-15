package jwtapi

import (
	"errors"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Foo string `json:"foo"`
}

func Parse(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("your-256-bit-secret"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return jwt.MapClaims{}, err
	}
}

func ValidateAlg(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return []byte("your-256-bit-secret"), nil
}

func ParseClaims(tokenString string) (Claims, error) {
	token, err := jwt.Parse(tokenString, ValidateAlg)
	if err != nil {
		return Claims{}, errors.New("Token error")
	}
	claims := token.Claims.(jwt.MapClaims)
	var c Claims
	c.Foo = claims["foo"].(string)
	return c, nil
}
