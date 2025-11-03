package permissions

// Simple systems store a single role; return as slice for uniformity.
import ur "golang-user-manager/internal/user/repository"

type ListUserRolesUseCase struct{ Repo ur.UserRepository }

func NewListUserRolesUseCase(r ur.UserRepository) ListUserRolesUseCase {
	return ListUserRolesUseCase{Repo: r}
}

func (uc ListUserRolesUseCase) Execute(userID uint) ([]string, error) {
	u, err := uc.Repo.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return []string{u.Role}, nil
}
