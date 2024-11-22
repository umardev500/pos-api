package pkg

import (
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
