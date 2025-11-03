package validation

import ur "golang-user-manager/internal/user/repository"

type CheckEmailUniquenessUseCase struct{ Repo ur.UserRepository }

func NewCheckEmailUniquenessUseCase(r ur.UserRepository) CheckEmailUniquenessUseCase {
	return CheckEmailUniquenessUseCase{Repo: r}
}

func (uc CheckEmailUniquenessUseCase) Execute(email string) (bool, error) {
	return uc.Repo.EmailExists(email)
}
