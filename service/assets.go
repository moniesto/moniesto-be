package service

import (
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/validation"
)

func (service *Service) GetConfigs() model.GetConfigsResponse {

	configs := model.GetConfigsResponse{
		ErrorCodes:  service.GetErrorCodes(),
		Validation:  service.GetValidationConfigs(),
		GeneralInfo: service.GetGeneralInfoConfigs(),
	}

	return configs
}

func (service *Service) GetGeneralInfoConfigs() model.GetGeneralInfoResponse {
	general_info := model.GetGeneralInfoResponse{
		OperationFeePercentage: service.config.OperationFeePercentage,
	}

	return general_info
}

func (service *Service) GetErrorCodes() clientError.ErrorMessagesType {
	return clientError.GetErrorCodes()
}

func (service *Service) GetValidationConfigs() model.GetValidationConfigsResponse {

	validation_configs := model.GetValidationConfigsResponse{
		EmailRegex:                      validation.EmailRegex,
		UsernameRegex:                   validation.UsernameRegex,
		MinFee:                          service.config.MinFee,
		MaxFee:                          service.config.MaxFee,
		MaxBioLenght:                    validation.MaxBioLength,
		MaxDescriptionLength:            validation.MaxDescriptionLength,
		MaxSubscriptionMessageLength:    validation.MaxSubscriptionMessageLength,
		MaxFullnameLength:               validation.MaxFullnameLength,
		MaxLocationLength:               validation.MaxLocationLength,
		MinPasswordLength:               validation.MinPasswordLength,
		MaxPasswordLength:               validation.MaxPasswordLength,
		MaxDurationDay:                  validation.MaxDurationDay,
		LongMaxTakeProfitMultiplierLong: validation.MaxTakeProfitMultiplierLong,
		ShortMaxStopMultiplier:          validation.MaxStopMultiplierShort,
	}

	return validation_configs
}
