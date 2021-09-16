package utils

import "encoding/hex"

// HexToBytes safely transforms a hexadecimal string to a byte slice.
func HexToBytes(hexStr string) []byte {
	bytes, err := hex.DecodeString(hexStr)
	PanicError(err)

	return bytes
}

// BytesToHex transforms a byte slice to a hexadecimal string.
func BytesToHex(bytes []byte) string {
	hexStr := hex.EncodeToString(bytes)

	return hexStr
}
