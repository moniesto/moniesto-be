package validation

import (
	"fmt"

	"net/mail"

	"github.com/moniesto/moniesto-be/util"
)

func Password(password string) error {
	if len(password) < util.ValidPasswordLength {
		return fmt.Errorf("length of password is less than %d", util.ValidPasswordLength)
	}

	return nil
}

func Email(email string) (string, error) {
	addr, err := mail.ParseAddress(email)

	if err != nil {
		return "", fmt.Errorf("email is not valid %s", email)
	}

	return addr.Address, nil
}
