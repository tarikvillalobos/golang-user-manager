package auth

import (
	"errors"

	ur "golang-user-manager/internal/user/repository"
	"golang-user-manager/pkg/hash"
)

type UpdatePasswordUseCase struct{ Repo ur.UserRepository }

func NewUpdatePasswordUseCase(r ur.UserRepository) UpdatePasswordUseCase {
	return UpdatePasswordUseCase{Repo: r}
}

func (uc UpdatePasswordUseCase) Execute(userID uint, newPass string) error {
	if len(newPass) < 6 {
		return errors.New("password too short")
	}
	u, err := uc.Repo.GetByID(userID)
	if err != nil {
		return err
	}
	h, err := hash.HashPassword(newPass)
	if err != nil {
		return err
	}
	u.Password = h
	return uc.Repo.Update(u)
}
