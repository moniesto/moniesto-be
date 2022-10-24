package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns the bcrypt hash of the password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("failed to password operation")
	}

	return string(hashedPassword), nil
}

// CheckPassword checks the provided password is correct or not
func CheckPassword(password, hashedPasword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPasword), []byte(password))
}
