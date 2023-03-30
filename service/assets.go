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
		MaxNameLength:                validation.MaxNameLength,
		MaxSurnameLength:             validation.MaxSurnameLength,
		MaxLocationLength:            validation.MaxLocationLength,
		PasswordLength:               validation.ValidPasswordLength,
	}

	return validation_configs
}
