package notifications

import "golang-user-manager/pkg/email"

type SendPasswordResetUseCase struct{ Mail email.Service }

func NewSendPasswordResetUseCase(m email.Service) SendPasswordResetUseCase {
	return SendPasswordResetUseCase{Mail: m}
}

func (uc SendPasswordResetUseCase) Execute(to, link string) error {
	return uc.Mail.Send(to, "Reset your password", "Clique para redefinir: "+link)
}
