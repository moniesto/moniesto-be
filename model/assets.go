package model

import "github.com/moniesto/moniesto-be/util/clientError"

type GetConfigsResponse struct {
	ErrorCodes clientError.ErrorMessagesType `json:"error_codes"`
	Validation GetValidationConfigsResponse  `json:"validation"`
}

type GetValidationConfigsResponse struct {
	EmailRegex                   string  `json:"email_regex"`
	UsernameRegex                string  `json:"username_regex"`
	MinFee                       float64 `json:"min_fee"`
	MaxBioLenght                 int     `json:"max_bio_lenght"`
	MaxDescriptionLength         int     `json:"max_description_length"`
	MaxSubscriptionMessageLength int     `json:"max_subscription_message_length"`
	PasswordLength               int     `json:"password_length"`
}
