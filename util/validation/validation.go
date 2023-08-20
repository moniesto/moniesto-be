package validation

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/moniesto/moniesto-be/config"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/util"
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

func Fullname(fullname string) error {
	if len(fullname) > MaxFullnameLength {
		return fmt.Errorf("fullname length is more than %d", MaxFullnameLength)
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

// Duration checks if is valid -> after now & less than 90 days
func Duration(duration time.Time) error {

	maxDayDate := util.Now().Add(time.Hour * 24 * time.Duration(MaxDurationDay))

	if util.Now().After(duration) || duration.After(maxDayDate) {
		return fmt.Errorf("invalid duration: more than %d day or assigned for past", MaxDurationDay)
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

		// max target can only be MaxTargetMultiplierLong times more than price
		if target3 > price*float64(MaxTargetMultiplierLong) {
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
		// stop can not be smaller than price & MaxStopMultiplierShort times more than price
		if stop < price || stop > price*float64(MaxStopMultiplierShort) {
			return fmt.Errorf(error_message)
		}
	} else {
		return fmt.Errorf(direction_error_message)
	}

	return nil
}

func Language(language string) error {
	if contains(supportedLanguages, language) {
		return nil
	}

	return fmt.Errorf("language is not supported: %s", language)
}

func BinanceID(binance_id string) error {

	if len(binance_id) > 0 {
		return nil
	}

	return fmt.Errorf("binance_id is not valid")

}

func UserIsAdmin(email string) bool {
	return contains(adminEmails, strings.ToLower(email))
}

func SubscriptionDateValue(number_of_months int) bool {
	if number_of_months >= 1 && number_of_months <= MaxSubscriptionMonth {
		return true
	}

	return false
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
