package service

import (
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/validation"
)

func (service *Service) GetConfigs() model.GetConfigsResponse {

	configs := model.GetConfigsResponse{
		ErrorCodes: service.GetErrorCodes(),
		Validation: service.GetValidationConfigs(),
	}

	return configs
}

func (service *Service) GetErrorCodes() clientError.ErrorMessagesType {
	return clientError.GetErrorCodes()
}

func (service *Service) GetValidationConfigs() model.GetValidationConfigsResponse {

	validation_configs := model.GetValidationConfigsResponse{
		EmailRegex:                   validation.EmailRegex,
		UsernameRegex:                validation.UsernameRegex,
		MinFee:                       service.config.MinFee,
		MaxBioLenght:                 validation.MaxBioLength,
		MaxDescriptionLength:         validation.MaxDescriptionLength,
		MaxSubscriptionMessageLength: validation.MaxSubscriptionMessageLength,
		MaxFullnameLength:            validation.MaxFullnameLength,
		MaxLocationLength:            validation.MaxLocationLength,
		PasswordLength:               validation.ValidPasswordLength,
		MaxDurationDay:               validation.MaxDurationDay,
		LongMaxTargetMultiplier:      validation.MaxTargetMultiplierLong,
		ShortMaxStopMultiplier:       validation.MaxStopMultiplierShort,
		// TODO: OPERATION_FEE_PERCENTAGE
	}

	return validation_configs
}
