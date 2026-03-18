package helper

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
)

func IsAuthorized(t string, k string) error {
	if t == "" {
		return errors.New("token cannot be empty")
	}

	if k == "" {
		return errors.New("key cannot be empty")
	}

	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(k), nil
	})

	if err != nil {
		Log(err, ErrorLevel)
		return err
	}

	if token.Valid {
		return nil
	} else {
		return errors.New("invalid token")
	}
}
