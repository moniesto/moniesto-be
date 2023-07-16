package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomEmail() string {
	return RandomString(8) + "@" + RandomString(5) + "." + RandomString(3)
}

func RandomUsername() string {
	return RandomString(10)
}

func RandomPassword() string {
	return RandomString(6)
}

func RandomFullname() string {
	return RandomString(6) + RandomString(6)
}
