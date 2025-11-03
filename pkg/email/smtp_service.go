package email

import "log"

type Service interface {
	Send(to, subject, body string) error
}

type DummyService struct{}

func (d DummyService) Send(to, subject, body string) error {
	log.Printf("[email] to=%s subject=%s body=%s", to, subject, body)
	return nil
}
