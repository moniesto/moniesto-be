package service

import "github.com/moniesto/moniesto-be/util/clientError"

func (service *Service) GetErrorCodes() clientError.ErrorMessagesType {
	return clientError.GetErrorCodes()
}
