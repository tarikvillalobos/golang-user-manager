package jwtutil

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("dev-secret-change-me")

type Claims struct {
	UserID uint   `json:"uid"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func Generate(userID uint, email, role string, ttl time.Duration) (string, error) {
	claims := &Claims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func Parse(tokenStr string) (*Claims, error) {
	t, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := t.Claims.(*Claims)
	if !ok || !t.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
