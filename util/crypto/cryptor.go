package crypto

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util/system"
)

// "https://api3.binance.com/api/v3/ticker/price",

var APIlinks []string = []string{
	"https://api3.binance.com/api/v3",
	"https://api2.binance.com/api/v3",
	"https://api1.binance.com/api/v3",
	"https://api.binance.com/api/v3",
}

var tickerURI = "/ticker/price"

// GetCurrencies get all currencies from the crypto API
func GetCurrencies() (model.GetCurrenciesAPIResponse, error) {
	var currencies model.GetCurrenciesAPIResponse

	client := resty.New()

	link_number := 0

	// check all APIs from the list
	for {
		api_link := APIlinks[link_number] + tickerURI

		resp, err := client.R().
			SetResult(&currencies).
			Get(api_link)

		// if fails, then check another API
		if err != nil || resp.StatusCode() >= 400 {
			if link_number+1 == len(APIlinks) { // no more new API
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

func GetCurrency(name string) (model.GetCurrencyAPIResponse, error) {
	var currency model.GetCurrencyAPIResponse

	client := resty.New()

	link_number := 0

	for {
		api_link := APIlinks[link_number] + tickerURI + "?symbol=" + name

		resp, err := client.R().
			SetResult(&currency).
			Get(api_link)

		// if fails, then check another API
		if err != nil || resp.StatusCode() >= 400 {
			if link_number+1 == len(APIlinks) { // no more new API
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
