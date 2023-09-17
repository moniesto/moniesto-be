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
		if len(username) == 0 || strings.util.Contains(username, " ") || util.Contains(InvalidUsernames, username) {
			return fmt.Errorf("username is not valid %s", username)
		}

		return nil
	*/

	usernameRegex := regexp.MustCompile(UsernameRegex)
	if !usernameRegex.MatchString(username) || util.Contains(InvalidUsernames, username) {
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

// TakeProfit checks is take profit valid [price > takeProfit > MaxTarget]
func TakeProfit(price, takeProfit float64, direction db.Direction) error {
	takeProfitErrorMessage := "take profit is not valid"
	directionErrorMessage := "direction is not valid"

	switch direction {
	case db.DirectionLong:
		if (takeProfit <= price) || (takeProfit > price*float64(MaxTakeProfitMultiplierLong)) {
			return fmt.Errorf(takeProfitErrorMessage)
		}
	case db.DirectionShort:
		if takeProfit >= price {
			return fmt.Errorf(takeProfitErrorMessage)
		}
	default:
		return fmt.Errorf(directionErrorMessage)
	}

	return nil
}

// Target checks targets are valid [price < target1 < target2 < target3 < takeProfit] => for long | or vice-versa for short
func Target(price float64, takeProfit float64, target1P, target2P, target3P *float64, direction db.Direction) error {
	targetErrorMessage := "targets are not valid"

	if (target3P != nil) && (target1P == nil || target2P == nil) { // tp3 is not nil, but one of tp1 or tp2 is nil
		return fmt.Errorf(targetErrorMessage)
	} else if target2P == nil && target3P != nil { // tp2 is nil, but tp3 is not
		return fmt.Errorf(targetErrorMessage)
	} else if target1P == nil && (target2P != nil || target3P != nil) { // tp1 is nil, but one of tp2 or tp3 is not nil
		return fmt.Errorf(targetErrorMessage)
	}

	targetsP := []*float64{target1P, target2P, target3P}
	targets := []float64{}

	for _, targetP := range targetsP {
		if targetP == nil {
			break
		}
		targets = append(targets, *targetP)
	}

	if len(targets) == 0 {
		return nil
	}

	var checkpoint float64 = 0
	for i, target := range targets {
		if i == 0 {
			checkpoint = target
		}

		switch direction {
		case db.DirectionLong:
			if !(price < target) || !(target < takeProfit) || (i != 0 && !(checkpoint < target)) {
				return fmt.Errorf(targetErrorMessage)
			}
		case db.DirectionShort:
			if !(target < price) || !(takeProfit < target) || (i != 0 && !(checkpoint > target)) {
				return fmt.Errorf(targetErrorMessage)
			}
		default:
			return fmt.Errorf(ERROR_DirectionNotValid)
		}

		checkpoint = target
	}

	return nil
}

// Stop checks stop is valid [stop < price]
func Stop(price, stop float64, leverage int32, direction db.Direction) error {
	error_message := "stop is not valid"

	if stop <= 0 {
		return fmt.Errorf(error_message)
	}

	var stopLimitPercentage float64 = float64(100 / float64(leverage))

	switch direction {
	case db.DirectionLong:
		if !(stop < price) || !(stop > (price - (price * (stopLimitPercentage / 100)))) {
			return fmt.Errorf(error_message)
		}
	case db.DirectionShort:
		if stop <= price || stop > price*float64(MaxStopMultiplierShort) || !(stop < (price + (price * (stopLimitPercentage / 100)))) {
			return fmt.Errorf(error_message)
		}
	default:
		return fmt.Errorf(ERROR_DirectionNotValid)
	}

	return nil
}

// MarketType checks provided market type is supported or not
func MarketType(marketType string) error {
	if util.Contains(supportedMarketTypes, marketType) {
		return nil
	}

	return fmt.Errorf("market type is not supported: %s", marketType)
}

// Leverage checks leverage is valid or not
func Leverage(leverage int32) error {
	if leverage >= MinLeverage && leverage <= MaxLeverage {
		return nil
	}

	return fmt.Errorf("leverage is not valid: %d", leverage)
}

// Language checks provided language is supported or not
func Language(language string) error {
	if util.Contains(supportedLanguages, language) {
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
	return util.Contains(adminEmails, strings.ToLower(email))
}

func SubscriptionDateValue(number_of_months int) bool {
	if number_of_months >= 1 && number_of_months <= MaxSubscriptionMonth {
		return true
	}

	return false
}
