package usecase

import sr "golang-user-manager/internal/subscription/repository"
import sd "golang-user-manager/internal/subscription/domain"

type CreateSubscriptionUseCase struct{ Repo sr.SubscriptionRepository }

func NewCreateSubscriptionUseCase(r sr.SubscriptionRepository) CreateSubscriptionUseCase {
	return CreateSubscriptionUseCase{Repo: r}
}

func (uc CreateSubscriptionUseCase) Execute(s *sd.Subscription) error { return uc.Repo.Create(s) }
