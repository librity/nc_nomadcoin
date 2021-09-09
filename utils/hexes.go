package utils

import "encoding/hex"

func HexToBytes(hexStr string) []byte {
	bytes, err := hex.DecodeString(hexStr)
	HandleError(err)

	return bytes
}

func BytesToHex(bytes []byte) string {
	hexStr := hex.EncodeToString(bytes)

	return hexStr
}
