package hash

import "golang.org/x/crypto/bcrypt"

func HashPassword(p string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	return string(b), err
}

func CheckPassword(hashed, plain string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain)) == nil
}
