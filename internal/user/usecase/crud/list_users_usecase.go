package crud

import ur "golang-user-manager/internal/user/repository"
import ud "golang-user-manager/internal/user/domain"

type ListUsersUseCase struct{ Repo ur.UserRepository }

func NewListUsersUseCase(r ur.UserRepository) ListUsersUseCase { return ListUsersUseCase{Repo: r} }

func (uc ListUsersUseCase) Execute() ([]ud.User, error) {
	return uc.Repo.List()
}
