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

func GenerateFloatSlice(n int) []float64 {
	floatSlice := make([]float64, 0, n)

	for i := 1; i <= n; i++ {
		floatSlice = append(floatSlice, float64(i))
	}

	return floatSlice
}

// GetRandomTime generate random time between 1 week and 2 months
func GetRandomTime() time.Time {
	// Get current time
	currentTime := Now()

	// Generate a random duration between 1 week and 3 months
	randomDuration := time.Duration(rand.Intn(8*7*24*3600)+1) * time.Second // 1 week to 2 months in seconds

	// Add the random duration to the current time
	randomTime := currentTime.Add(randomDuration)

	// Ensure the generated time is at least 1 week from the current time
	minimumTime := currentTime.Add(7 * 24 * time.Hour)
	if randomTime.Before(minimumTime) {
		randomTime = minimumTime
	}

	return randomTime
}
