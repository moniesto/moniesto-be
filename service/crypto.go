package service

import (
	"net/http"
	"strings"

	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util/clientError"
	"github.com/moniesto/moniesto-be/util/crypto"
)

func (service *Service) GetCurrenciesWithName(name string) ([]model.Currency, error) {
	// STEP: get all currencies
	currencies, err := crypto.GetCurrencies()
	if err != nil || len(currencies) == 0 {
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
