package crud

import (
	"errors"
	"regexp"

	ud "golang-user-manager/internal/user/domain"
	ur "golang-user-manager/internal/user/repository"
	"golang-user-manager/pkg/hash"
)

type CreateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type CreateUserUseCase struct{ Repo ur.UserRepository }

func NewCreateUserUseCase(r ur.UserRepository) CreateUserUseCase { return CreateUserUseCase{Repo: r} }

func (uc CreateUserUseCase) Execute(in CreateUserDTO) (*ud.User, error) {
	if !validName(in.Name) {
		return nil, errors.New("invalid name")
	}
	if !validEmail(in.Email) {
		return nil, errors.New("invalid email")
	}
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
	if in.Role != "" {
		u.Role = in.Role
	}
	return u, uc.Repo.Create(u)
}

var emailRx = regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)

func validEmail(s string) bool { return emailRx.MatchString(s) }
func validName(s string) bool  { return len(s) >= 2 }
