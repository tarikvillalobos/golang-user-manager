package validation

import ur "golang-user-manager/internal/user/repository"

type CheckUserExistsUseCase struct{ Repo ur.UserRepository }

func NewCheckUserExistsUseCase(r ur.UserRepository) CheckUserExistsUseCase {
	return CheckUserExistsUseCase{Repo: r}
}

func (uc CheckUserExistsUseCase) Execute(id uint) (bool, error) { return uc.Repo.Exists(id) }
