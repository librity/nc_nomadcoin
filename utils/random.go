package utils

import (
	"math/rand"
	"time"
)

const (
	alphanum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// RandomString generates a rando alpha-numeric string.
func RandomString(length int) string {
	SeedRNG()

	randStr := make([]byte, length)
	for i := range randStr {
		random := rand.Intn(len(alphanum))
		randStr[i] = alphanum[random]
	}

	return string(randStr)
}

// SeedRNG seeds the Random Number Generator with the current time.
func SeedRNG() {
	rand.Seed(time.Now().UnixNano())
}
