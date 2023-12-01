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
