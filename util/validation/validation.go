package validation

import (
	"fmt"
	"strings"

	"net/mail"

	"github.com/moniesto/moniesto-be/config"
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
	if len(username) == 0 || strings.Contains(username, " ") || contains(InvalidUsernames, username) {
		return fmt.Errorf("username is not valid %s", username)
	}

	return nil
}

// Fee checks the fee is valid
func Fee(fee float64, config config.Config) error {
	if fee < config.MinFee {
		return fmt.Errorf("fee is not valid %f", fee)
	}

	return nil
}

// Bio checks the bio is valid
func Bio(bio string, config config.Config) error {
	if len(bio) > config.MaxBioLenght {
		return fmt.Errorf("bio is not valid %s", bio)
	}

	return nil
}

// Description checks the description is valid
func Description(description string, config config.Config) error {
	if len(description) > config.MaxDescriptionLength {
		return fmt.Errorf("description is not valid %s", description)
	}

	return nil
}

// SubscriptionMessage checks the message is valid
func SubscriptionMessage(message string, config config.Config) error {
	if len(message) > config.MaxSubscriptionMessageLength {
		return fmt.Errorf("message is not valid %s", message)
	}

	return nil
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
