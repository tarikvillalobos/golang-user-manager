package usecase

import (
	"errors"

	ur "golang-user-manager/internal/user/repository"
	"golang-user-manager/pkg/hash"
	"golang-user-manager/pkg/jwt"
)

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserUseCase struct{ Repo ur.UserRepository }

func NewLoginUserUseCase(r ur.UserRepository) LoginUserUseCase { return LoginUserUseCase{Repo: r} }

func (uc LoginUserUseCase) Execute(in LoginDTO) (string, error) {
	u, err := uc.Repo.GetByEmail(in.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	if !hash.CheckPassword(u.Password, in.Password) {
		return "", errors.New("invalid credentials")
	}
	tok, err := jwtutil.Generate(u.ID, u.Email, u.Role, 24*60*60*1e9)
	return tok, err
}
