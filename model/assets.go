package model

import "github.com/moniesto/moniesto-be/util/clientError"

type GetConfigsResponse struct {
	Validation  GetValidationConfigsResponse  `json:"validation"`
	ErrorCodes  clientError.ErrorMessagesType `json:"error_codes"`
	GeneralInfo GetGeneralInfoResponse        `json:"general_info"`
}

type GetValidationConfigsResponse struct {
	EmailRegex                   string  `json:"email_regex"`
	UsernameRegex                string  `json:"username_regex"`
	MinFee                       float64 `json:"min_fee"`
	MaxBioLenght                 int     `json:"max_bio_lenght"`
	MaxDescriptionLength         int     `json:"max_description_length"`
	MaxSubscriptionMessageLength int     `json:"max_subscription_message_length"`
	MaxFullnameLength            int     `json:"max_fullname_length"`
	MaxLocationLength            int     `json:"max_location_length"`
	PasswordLength               int     `json:"password_length"`

	MaxDurationDay          int `json:"max_duration_day"`
	LongMaxTargetMultiplier int `json:"long_max_target_multiplier"`
	ShortMaxStopMultiplier  int `json:"short_max_stop_multiplier"`
}

type GetGeneralInfoResponse struct {
	OperationFeePercentage float64 `json:"operation_fee_percentage"`
}
