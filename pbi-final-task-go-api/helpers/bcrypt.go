package helpers

import "golang.org/x/crypto/bcrypt"

func GenerateFromPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 10)
}

func CompareHashAndPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
