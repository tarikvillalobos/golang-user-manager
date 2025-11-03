package notifications

import "golang-user-manager/pkg/logger"

type NotifyUserDeletedUseCase struct{}

func (NotifyUserDeletedUseCase) Execute(userID uint) {
	logger.Info("user.deleted", "userID", userID)
}
