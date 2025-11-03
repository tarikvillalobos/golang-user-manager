package usecase

import (
	"errors"

	ud "golang-user-manager/internal/user/domain"
	ur "golang-user-manager/internal/user/repository"
	"golang-user-manager/pkg/hash"
)

type RegisterUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserUseCase struct{ Repo ur.UserRepository }

func NewRegisterUserUseCase(r ur.UserRepository) RegisterUserUseCase {
	return RegisterUserUseCase{Repo: r}
}

func (uc RegisterUserUseCase) Execute(in RegisterUserDTO) (*ud.User, error) {
	exists, err := uc.Repo.EmailExists(in.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("email already in use")
	}
	h, err := hash.HashPassword(in.Password)
	if err != nil {
		return nil, err
	}
	u := &ud.User{Name: in.Name, Email: in.Email, Password: h}
	return u, uc.Repo.Create(u)
}
