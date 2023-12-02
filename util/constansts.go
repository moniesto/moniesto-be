package util

import (
	"math/rand"
)

var systemMoniests = []string{
	"astrader",
	"parvin",
	"cryptofamily",
	"lowriskman",
	"onebull",
	"booster_group",
	"cryptomonster",
	"astrader",

	"kainpainter",
	"layneschmitz",
	"sukrukara",
	"mustafayuksel",
	"danabarnhill",
	"johnnamunn",
	"lizetpoirier8",
	"journeyfarr2",
	"gregoriobreen",
	"rosalindaher",
	"milesricketts",
	"yusufcan7",
}

var TwitterPostContent string = `🚀 Check out #$trader_username's #$coin_name analysis! 📈
	
Earned extra $$pnl on a $1000 investment 💰

https://moniesto.com/$trader_username

#$coin_name #BTC #Binance #CryptoInsights #Crypto #CryptoAnalyze
`

func GetSystemMoniests() []string {
	return systemMoniests
}

func RandomMoniestUsername(usernames []string) string {
	return Random(usernames)
}

func Random[T comparable](slice []T) T {
	randomIndex := rand.Intn(len(slice))
	return slice[randomIndex]
}
