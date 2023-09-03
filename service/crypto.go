package service

import (
	"net/http"
	"strings"

	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/crypto"
	"github.com/moniesto/moniesto-be/util/system"
)

func (service *Service) GetCurrenciesWithName(name string, marketType string) ([]model.Currency, error) {
	// STEP: get all currencies
	currencies, err := crypto.GetCurrencies(marketType)
	if err != nil || len(currencies) == 0 {
		system.LogError("server error on get currencies", err.Error())
		return []model.Currency{}, clientError.CreateError(http.StatusInternalServerError, clientError.Crypto_GetCurrenciesFromAPI_ServerError)
	}

	// STEP: filter by name
	var filteredCurrencies []model.Currency = []model.Currency{}
	for _, currency := range currencies {
		if strings.Contains(strings.ToLower(currency.Symbol), strings.ToLower(name)) {
			filteredCurrencies = append(filteredCurrencies, model.Currency{
				Currency: currency.Symbol,
				Price:    currency.Price,
			})
		}
	}

	return filteredCurrencies, nil
}

func (service *Service) GetCurrency(name string, marketType string) (model.Currency, error) {

	// STEP: get single currency
	currency, err := crypto.GetCurrency(name, marketType)
	if err != nil {
		system.LogError("server error on get currency", err.Error())
		return model.Currency{}, clientError.CreateError(http.StatusInternalServerError, clientError.Crypto_GetCurrencyFromAPI_ServerError)
	}

	return model.Currency{
		Currency: currency.Symbol,
		Price:    currency.Price,
	}, nil
}
