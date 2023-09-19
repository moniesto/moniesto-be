package crypto

import db "github.com/moniesto/moniesto-be/db/sqlc"

var spotApiLinks []string = []string{
	"https://api3.binance.com/api/v3",
	"https://api2.binance.com/api/v3",
	"https://api1.binance.com/api/v3",
	"https://api.binance.com/api/v3",
}

var futuresApiLinks []string = []string{
	"https://fapi.binance.com/fapi/v1",
}

var MARKETS map[string][]string = map[string][]string{
	string(db.PostCryptoMarketTypeSpot):    spotApiLinks,
	string(db.PostCryptoMarketTypeFutures): futuresApiLinks,
}

const tickerURI = "/ticker/price"
const exchangeInfoURI = "/exchangeInfo"
const historyURI = "/klines"

const INTERVAL_1second = "1s"
