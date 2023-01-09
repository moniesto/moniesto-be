package service

import (
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/validation"
)

func (service *Service) GetErrorCodes() clientError.ErrorMessagesType {
	return clientError.GetErrorCodes()
}

func (service *Service) GetConfigs() model.GetConfigsResponse {

	configs := model.GetConfigsResponse{
		EmailRegex:                   validation.EmailRegex,
		UsernameRegex:                validation.UsernameRegex,
		MinFee:                       service.config.MinFee,
		MaxBioLenght:                 service.config.MaxBioLenght,
		MaxDescriptionLength:         service.config.MaxDescriptionLength,
		MaxSubscriptionMessageLength: service.config.MaxSubscriptionMessageLength,
		PasswordLength:               validation.ValidPasswordLength,
	}

	return configs
}
