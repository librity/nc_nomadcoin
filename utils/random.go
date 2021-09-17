package utils

import (
	"math/rand"
	"time"
)

const (
	alphanum     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	lowercaseHex = "0123456789abcdef"
)

// RandomHash generates a random lowercase hex hash.
func RandomHash() string {
	return RandomString(lowercaseHex, 64)
}

// RandomAlphanum generates a random alpha-numeric string.
func RandomAlphanum(length int) string {
	return RandomString(alphanum, length)
}

// RandomString generates a random string.
func RandomString(cypher string, length int) string {
	SeedRNG()

	randStr := make([]byte, length)
	for i := range randStr {
		random := rand.Intn(len(cypher))
		randStr[i] = cypher[random]
	}

	return string(randStr)
}

// SeedRNG seeds the Random Number Generator with the current time.
func SeedRNG() {
	rand.Seed(time.Now().UnixNano())
}
