package utils

import (
	"crypto/sha256"
	"fmt"
)

func HexHash(data interface{}) string {
	defaultFormat := fmt.Sprintf("%v", data)

	return HexHashStr(defaultFormat)
}

func HexHashStr(data string) string {
	dataBytes := []byte(data)
	rawHash := sha256.Sum256(dataBytes)
	hexHash := fmt.Sprintf("%x", rawHash)

	return hexHash
}
