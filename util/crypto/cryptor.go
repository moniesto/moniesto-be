package crypto

import (
	"github.com/go-resty/resty/v2"
	"github.com/moniesto/moniesto-be/model"
)

var APIlinks []string = []string{
	"https://api3.binance.com/api/v3/ticker/price",
	"https://api2.binance.com/api/v3/ticker/price",
	"https://api1.binance.com/api/v3/ticker/price",
	"https://api.binance.com/api/v3/ticker/price",
}

// GetCurrencies get all currencies from the crypto API
func GetCurrencies() (model.GetCurrenciesAPIResponse, error) {
	var currencies model.GetCurrenciesAPIResponse

	client := resty.New()

	link_number := 0

	// check all APIs from the list
	for {
		resp, err := client.R().
			SetResult(&currencies).
			Get(APIlinks[link_number])

		// if fails, then check another API
		if err != nil || resp.StatusCode() >= 400 {
			if link_number+1 == len(APIlinks) { // no more new API
				return model.GetCurrenciesAPIResponse{}, err
			} else {
				link_number = link_number + 1
				continue
			}
		}

		return currencies, nil
	}
}
