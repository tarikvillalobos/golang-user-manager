package permissions

import ur "golang-user-manager/internal/user/repository"

type AssignRoleToUserUseCase struct{ Repo ur.UserRepository }

func NewAssignRoleToUserUseCase(r ur.UserRepository) AssignRoleToUserUseCase {
	return AssignRoleToUserUseCase{Repo: r}
}

func (uc AssignRoleToUserUseCase) Execute(userID uint, role string) error {
	u, err := uc.Repo.GetByID(userID)
	if err != nil {
		return err
	}
	u.Role = role
	return uc.Repo.Update(u)
}
