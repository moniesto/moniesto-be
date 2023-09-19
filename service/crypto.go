package service

import (
	"net/http"
	"strings"

	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util"
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

	// STEP: filter quarterly currencies [like: BTCUSDT_230929]
	if marketType == string(db.PostCryptoMarketTypeFutures) {
		filteredCurrencies = filterQuarterlyCurrencies(&filteredCurrencies)
	}

	// STEP: filter only active currencies
	activeCurrencies, err := filterInActiveCurrencies(&filteredCurrencies)
	if err != nil {
		return []model.Currency{}, clientError.CreateError(http.StatusInternalServerError, clientError.Crypto_GetCurrenciesFromAPI_ServerError)
	}

	return activeCurrencies, nil
}

func (service *Service) GetCurrency(name string, marketType string) (model.Currency, error) {
	// STEP: get single currency
	currencyResponse, err := crypto.GetCurrency(name, marketType)
	if err != nil {
		system.LogError("server error on get currency", err.Error())
		return model.Currency{}, clientError.CreateError(http.StatusInternalServerError, clientError.Crypto_GetCurrencyFromAPI_ServerError)
	}

	filteredCurrencies := []model.Currency{{
		Currency: currencyResponse.Symbol,
		Price:    currencyResponse.Price,
	}}

	// STEP: filter quarterly currencies [like: BTCUSDT_230929]
	if marketType == string(db.PostCryptoMarketTypeFutures) {
		filteredCurrencies = filterQuarterlyCurrencies(&filteredCurrencies)
	}

	// STEP: filter only active currencies
	activeCurrencies, err := filterInActiveCurrencies(&filteredCurrencies)
	if err != nil {
		return model.Currency{}, clientError.CreateError(http.StatusInternalServerError, clientError.Crypto_GetCurrencyFromAPI_ServerError)
	}

	if len(activeCurrencies) != 1 {
		system.LogError("no currency left after filtering quarterly")
		return model.Currency{}, clientError.CreateError(http.StatusInternalServerError, clientError.Crypto_GetCurrencyFromAPI_ServerError)
	}

	return activeCurrencies[0], nil
}

// filterQuarterlyCurrencies return only currencies without quarterly values
func filterQuarterlyCurrencies(currencies *[]model.Currency) []model.Currency {
	filteredCurrencies := []model.Currency{}

	for _, currency := range *currencies {
		if !strings.Contains(currency.Currency, "_") {
			filteredCurrencies = append(filteredCurrencies, currency)
		}
	}

	return filteredCurrencies
}

// filterInActiveCurrencies returns only active currencies
func filterInActiveCurrencies(currencies *[]model.Currency) ([]model.Currency, error) {
	activeCurrencies, err := crypto.GetActiveCurrencies()
	if err != nil {
		return nil, err
	}

	filteredCurrencies := []model.Currency{}

	for _, currency := range *currencies {
		if util.Contains(activeCurrencies, currency.Currency) {
			filteredCurrencies = append(filteredCurrencies, currency)
		}
	}

	return filteredCurrencies, nil
}
