package pkg

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(claims jwt.MapClaims, secret string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func ValidateJWT(tokenString string, secret string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return &claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
