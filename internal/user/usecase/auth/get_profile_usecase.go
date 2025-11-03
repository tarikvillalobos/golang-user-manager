package auth

import ud "golang-user-manager/internal/user/domain"
import ur "golang-user-manager/internal/user/repository"

type GetProfileUseCase struct{ Repo ur.UserRepository }

func NewGetProfileUseCase(r ur.UserRepository) GetProfileUseCase { return GetProfileUseCase{Repo: r} }

func (uc GetProfileUseCase) Execute(userID uint) (*ud.User, error) { return uc.Repo.GetByID(userID) }
