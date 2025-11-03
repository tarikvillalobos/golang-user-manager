package usecase

import (
	"golang-user-manager/pkg/jwt"
	"time"
)

type RefreshTokenUseCase struct{}

func (RefreshTokenUseCase) Execute(token string) (string, error) {
	claims, err := jwtutil.Parse(token)
	if err != nil {
		return "", err
	}
	return jwtutil.Generate(claims.UserID, claims.Email, claims.Role, 24*time.Hour)
}
