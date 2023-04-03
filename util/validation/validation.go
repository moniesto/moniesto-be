package validation

import (
	"fmt"
	"regexp"

	"github.com/moniesto/moniesto-be/config"
	db "github.com/moniesto/moniesto-be/db/sqlc"
)

// Password checks the password is valid
func Password(password string) error {
	if len(password) < ValidPasswordLength {
		return fmt.Errorf("length of password is less than %d", ValidPasswordLength)
	}

	return nil
}

// Email checks the email is valid
func Email(email string) (string, error) {
	// IMPLEMENTATION 1
	/*
		addr, err := mail.ParseAddress(email)

		if err != nil {
			return "", fmt.Errorf("email is not valid %s", email)
		}

		return addr.Address, nil
	*/

	// IMPLEMENTATION 2 [updated with regex logic]
	emailRegex := regexp.MustCompile(EmailRegex)
	if !emailRegex.MatchString(email) {
		return "", fmt.Errorf("email is not valid %s", email)
	}

	return email, nil
}

// Username checks the username is valid
func Username(username string) error {
	// IMPLEMENTATION 1
	/*
		if len(username) == 0 || strings.Contains(username, " ") || contains(InvalidUsernames, username) {
			return fmt.Errorf("username is not valid %s", username)
		}

		return nil
	*/

	usernameRegex := regexp.MustCompile(UsernameRegex)
	if !usernameRegex.MatchString(username) || contains(InvalidUsernames, username) {
		return fmt.Errorf("username is not valid %s", username)
	}

	return nil
}

func Location(location string) error {
	if len(location) > MaxLocationLength {
		return fmt.Errorf("location length is more than %d", MaxLocationLength)
	}

	return nil
}

func Name(name string) error {
	if len(name) > MaxNameLength {
		return fmt.Errorf("name length is more than %d", MaxNameLength)
	}

	return nil
}

func Surname(surname string) error {
	if len(surname) > MaxSurnameLength {
		return fmt.Errorf("surname length is more than %d", MaxSurnameLength)
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
	if len(bio) > MaxBioLength {
		return fmt.Errorf("bio is not valid %s", bio)
	}

	return nil
}

// Description checks the description is valid
func Description(description string, config config.Config) error {
	if len(description) > MaxDescriptionLength {
		return fmt.Errorf("description is not valid %s", description)
	}

	return nil
}

// SubscriptionMessage checks the message is valid
func SubscriptionMessage(message string, config config.Config) error {
	if len(message) > MaxSubscriptionMessageLength {
		return fmt.Errorf("message is not valid %s", message)
	}

	return nil
}

// Target checks targets are valid [price < target1 < target2 < target3]
func Target(price, target1, target2, target3 float64, direction db.EntryPosition) error {
	target_error_message := "targets are not valid"
	direction_error_message := "direction is not valid"

	if target1 <= 0 || target2 <= 0 || target3 <= 0 {
		return fmt.Errorf(target_error_message)
	}

	if direction == db.EntryPositionLong {
		if !(price < target1) || !(target1 < target2) || !(target2 < target3) {
			return fmt.Errorf(target_error_message)
		}
	} else if direction == db.EntryPositionShort {
		if !(target3 < target2) || !(target2 < target1) || !(target1 < price) {
			return fmt.Errorf(target_error_message)
		}
	} else {
		return fmt.Errorf(direction_error_message)
	}

	return nil
}

// Stop checks stop is valid [stop < price]
func Stop(price, stop float64, direction db.EntryPosition) error {
	error_message := "stop is not valid"
	direction_error_message := "direction is not valid"

	if stop <= 0 {
		return fmt.Errorf(error_message)
	}

	if direction == db.EntryPositionLong {
		if !(stop < price) {
			return fmt.Errorf(error_message)
		}
	} else if direction == db.EntryPositionShort {
		if !(stop > price) {
			return fmt.Errorf(error_message)
		}
	} else {
		return fmt.Errorf(direction_error_message)
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
