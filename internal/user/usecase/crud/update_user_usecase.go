package crud

import (
	"errors"
	"regexp"

	ud "golang-user-manager/internal/user/domain"
	ur "golang-user-manager/internal/user/repository"
)

type UpdateUserDTO struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type UpdateUserUseCase struct{ Repo ur.UserRepository }

func NewUpdateUserUseCase(r ur.UserRepository) UpdateUserUseCase { return UpdateUserUseCase{Repo: r} }

func (uc UpdateUserUseCase) Execute(in UpdateUserDTO) (*ud.User, error) {
	u, err := uc.Repo.GetByID(in.ID)
	if err != nil {
		return nil, err
	}
	if in.Name != "" && len(in.Name) < 2 {
		return nil, errors.New("invalid name")
	}
	if in.Email != "" {
		if !regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`).MatchString(in.Email) {
			return nil, errors.New("invalid email")
		}
		// check uniqueness if changing email
		if in.Email != u.Email {
			exists, err := uc.Repo.EmailExists(in.Email)
			if err != nil {
				return nil, err
			}
			if exists {
				return nil, errors.New("email already in use")
			}
		}
		u.Email = in.Email
	}
	if in.Name != "" {
		u.Name = in.Name
	}
	if in.Role != "" {
		u.Role = in.Role
	}
	return u, uc.Repo.Update(u)
}
