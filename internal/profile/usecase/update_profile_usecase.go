package usecase

import pr "golang-user-manager/internal/profile/repository"
import pd "golang-user-manager/internal/profile/domain"

type UpdateProfileUseCase struct{ Repo pr.ProfileRepository }

func NewUpdateProfileUseCase(r pr.ProfileRepository) UpdateProfileUseCase {
	return UpdateProfileUseCase{Repo: r}
}

func (uc UpdateProfileUseCase) Execute(p *pd.Profile) error { return uc.Repo.Update(p) }
