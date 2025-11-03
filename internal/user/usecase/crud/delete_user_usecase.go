package crud

import ur "golang-user-manager/internal/user/repository"

type DeleteUserUseCase struct{ Repo ur.UserRepository }

func NewDeleteUserUseCase(r ur.UserRepository) DeleteUserUseCase { return DeleteUserUseCase{Repo: r} }

func (uc DeleteUserUseCase) Execute(id uint) error {
	return uc.Repo.Delete(id)
}
