package utils

import (
	"errors"
	"fmt"
	"math/big"
)

var (
	ErrBigIntBadBytes = errors.New("byte slice length should equal 32")
)

func BigIntsToHex(a, b *big.Int) string {
	bytes := BigIntsToBytes(a, b)
	hex := fmt.Sprintf("%x", bytes)

	return hex
}

func BigIntsToBytes(a, b *big.Int) []byte {
	bytes := append(a.Bytes(), b.Bytes()...)

	return bytes
}

func BigIntsFromHex(hex string) (*big.Int, *big.Int) {
	bytes := HexToBytes(hex)
	halfLength := len(bytes) / 2
	aBytes := bytes[:halfLength]
	bBytes := bytes[halfLength:]

	a, b := BigIntsFromBytes(aBytes, bBytes)
	return a, b
}

func BigIntsFromBytes(aBytes, bBytes []byte) (*big.Int, *big.Int) {
	a, err := BytesToBigInt(aBytes)
	HandleError(err)
	b, err := BytesToBigInt(bBytes)
	HandleError(err)

	return a, b
}

func BytesToBigInt(bytes []byte) (*big.Int, error) {
	if len(bytes) != 32 {
		return nil, ErrBigIntBadBytes
	}

	bi := new(big.Int).SetBytes(bytes)
	return bi, nil
}
