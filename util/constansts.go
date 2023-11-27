package util

import (
	"math/rand"
)

var systemMoniests = []string{
	"astrader",
	"parvin",
	"onebull",
	"booster_group",
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
