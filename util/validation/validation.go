package validation

import (
	"fmt"
	"strings"

	"net/mail"

	"github.com/moniesto/moniesto-be/util"
)

// Password checks the password is valid
func Password(password string) error {
	if len(password) < util.ValidPasswordLength {
		return fmt.Errorf("length of password is less than %d", util.ValidPasswordLength)
	}

	return nil
}

// Email checks the email is valid
func Email(email string) (string, error) {
	addr, err := mail.ParseAddress(email)

	if err != nil {
		return "", fmt.Errorf("email is not valid %s", email)
	}

	return addr.Address, nil
}

// Username checks the username is valid
func Username(username string) error {
	if len(username) == 0 || strings.Contains(username, " ") {
		return fmt.Errorf("username is not valid %s", username)
	}

	return nil
}
