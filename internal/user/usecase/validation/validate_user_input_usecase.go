package validation

import "regexp"

type ValidateUserInputUseCase struct{}

func (ValidateUserInputUseCase) NameOK(name string) bool { return len(name) >= 2 }
func (ValidateUserInputUseCase) EmailOK(email string) bool {
	return regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`).MatchString(email)
}
