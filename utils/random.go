package utils

import (
	"math/rand"
	"time"
)

const (
	alphanum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func RandomString(length int) string {
	SeedRandom()

	randStr := make([]byte, length)
	for i := range randStr {
		random := rand.Intn(len(alphanum))
		randStr[i] = alphanum[random]
	}

	return string(randStr)
}

func SeedRandom() {
	rand.Seed(time.Now().UnixNano())
}
