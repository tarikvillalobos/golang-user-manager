package crud

import ur "golang-user-manager/internal/user/repository"
import ud "golang-user-manager/internal/user/domain"

type GetUserByIDUseCase struct{ Repo ur.UserRepository }

func NewGetUserByIDUseCase(r ur.UserRepository) GetUserByIDUseCase {
	return GetUserByIDUseCase{Repo: r}
}

func (uc GetUserByIDUseCase) Execute(id uint) (*ud.User, error) {
	return uc.Repo.GetByID(id)
}
