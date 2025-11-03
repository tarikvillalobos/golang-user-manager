package permissions

import ur "golang-user-manager/internal/user/repository"

type RevokeRoleFromUserUseCase struct{ Repo ur.UserRepository }

func NewRevokeRoleFromUserUseCase(r ur.UserRepository) RevokeRoleFromUserUseCase {
	return RevokeRoleFromUserUseCase{Repo: r}
}

func (uc RevokeRoleFromUserUseCase) Execute(userID uint) error {
	u, err := uc.Repo.GetByID(userID)
	if err != nil {
		return err
	}
	u.Role = "user"
	return uc.Repo.Update(u)
}
