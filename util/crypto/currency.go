package crypto

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util/system"
)

// GetCurrencies get all currencies from the crypto API
func GetCurrencies(marketType string) (model.GetCurrenciesAPIResponse, error) {
	var currencies model.GetCurrenciesAPIResponse

	client := resty.New()

	apiLinks, ok := MARKETS[marketType]
	if !ok {
		return model.GetCurrenciesAPIResponse{}, fmt.Errorf("market type is not supported: %s", marketType)
	}

	link_number := 0
	// check all APIs from the list
	for {
		api_link := apiLinks[link_number] + tickerURI

		resp, err := client.R().
			SetResult(&currencies).
			Get(api_link)

		// if fails, then check another API
		if err != nil || resp.StatusCode() >= 400 {
			if link_number+1 == len(apiLinks) { // no more new API
				system.LogError("GetCurrencies: binance all APIs are unavaliable")
				return model.GetCurrenciesAPIResponse{}, fmt.Errorf("binance API error")
			} else {
				link_number = link_number + 1
				continue
			}
		}

		return currencies, nil
	}
}

func GetCurrency(name string, marketType string) (model.GetCurrencyAPIResponse, error) {
	var currency model.GetCurrencyAPIResponse

	client := resty.New()

	apiLinks, ok := MARKETS[marketType]
	if !ok {
		return model.GetCurrencyAPIResponse{}, fmt.Errorf("market type is not supported: %s", marketType)
	}

	link_number := 0
	for {
		api_link := apiLinks[link_number] + tickerURI + "?symbol=" + name

		resp, err := client.R().
			SetResult(&currency).
			Get(api_link)

		// if fails, then check another API
		if err != nil || resp.StatusCode() >= 400 {
			if link_number+1 == len(apiLinks) { // no more new API
				system.LogError("GetCurrency: binance all APIs are unavaliable")
				return model.GetCurrencyAPIResponse{}, fmt.Errorf("binance API error")
			} else {
				link_number = link_number + 1
				continue
			}
		}

		return currency, nil
	}
}
