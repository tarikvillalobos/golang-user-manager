package notifications

import "golang-user-manager/pkg/logger"

type LogUserActivityUseCase struct{}

func (LogUserActivityUseCase) Execute(action string, kv ...interface{}) {
	logger.Info("user.activity "+action, kv...)
}
