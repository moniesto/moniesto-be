package crypto

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util/system"
)

func GetHistories(symbol string, marketType string, inverval string, start_UTC_TS, end_UTC_TS int64, limit int) ([]model.History, error) {
	var histories []model.History

	client := resty.New()

	apiLinks, ok := MARKETS[string(marketType)]
	if !ok {
		return nil, fmt.Errorf("market type is not supported: %s", marketType)
	}

	link_number := 0
	// check all APIs from the list
	for {
		uriQueries := fmt.Sprintf("?symbol=%s&interval=%s&startTime=%d&endTime=%d&limit=%d", symbol, inverval, start_UTC_TS, end_UTC_TS, limit)

		api_link := apiLinks[link_number] + historyURI + uriQueries

		resp, err := client.R().SetResult(&histories).Get(api_link)

		// if fails, then check another API
		if err != nil || resp.StatusCode() >= 400 {
			if link_number+1 == len(apiLinks) { // no more new API
				system.LogError("GetHistories: binance all APIs are unavaliable")
				return nil, fmt.Errorf("binance API error")
			} else {
				link_number = link_number + 1
				continue
			}
		}

		return histories, nil
	}
}
