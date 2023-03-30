package model

import "github.com/moniesto/moniesto-be/util/clientError"

type GetConfigsResponse struct {
	Validation GetValidationConfigsResponse  `json:"validation"`
	ErrorCodes clientError.ErrorMessagesType `json:"error_codes"`
}

type GetValidationConfigsResponse struct {
	EmailRegex                   string  `json:"email_regex"`
	UsernameRegex                string  `json:"username_regex"`
	MinFee                       float64 `json:"min_fee"`
	MaxBioLenght                 int     `json:"max_bio_lenght"`
	MaxDescriptionLength         int     `json:"max_description_length"`
	MaxSubscriptionMessageLength int     `json:"max_subscription_message_length"`
	MaxNameLength                int     `json:"max_name_length"`
	MaxSurnameLength             int     `json:"max_surname_length"`
	MaxLocationLength            int     `json:"max_location_length"`
	PasswordLength               int     `json:"password_length"`
}
