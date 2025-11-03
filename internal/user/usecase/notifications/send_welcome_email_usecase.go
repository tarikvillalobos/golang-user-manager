package notifications

import "golang-user-manager/pkg/email"

type SendWelcomeEmailUseCase struct{ Mail email.Service }

func NewSendWelcomeEmailUseCase(m email.Service) SendWelcomeEmailUseCase {
	return SendWelcomeEmailUseCase{Mail: m}
}

func (uc SendWelcomeEmailUseCase) Execute(to string) error {
	return uc.Mail.Send(to, "Welcome", "Bem-vindo à plataforma!")
}
