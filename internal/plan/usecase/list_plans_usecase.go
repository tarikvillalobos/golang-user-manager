package usecase

import pr "golang-user-manager/internal/plan/repository"
import pd "golang-user-manager/internal/plan/domain"

type ListPlansUseCase struct{ Repo pr.PlanRepository }

func NewListPlansUseCase(r pr.PlanRepository) ListPlansUseCase { return ListPlansUseCase{Repo: r} }

func (uc ListPlansUseCase) Execute() ([]pd.Plan, error) { return uc.Repo.List() }
