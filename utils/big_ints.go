package utils

import (
	"errors"
	"fmt"
	"math/big"
)

var (
	ErrBigIntBadBytes = errors.New("byte slice length should equal 32")
)

// BigIntsToHex encodes two big.Int pointers as a hexdecimal string.
func BigIntsToHex(a, b *big.Int) string {
	bytes := BigIntsToBytes(a, b)
	hex := fmt.Sprintf("%x", bytes)

	return hex
}

// BigIntsToBytes encodes two big.Int pointers as a byte slice.
func BigIntsToBytes(a, b *big.Int) []byte {
	bytes := append(a.Bytes(), b.Bytes()...)

	return bytes
}

// BigIntsFromHex decodes two big.Int pointers from a hexdecimal string.
func BigIntsFromHex(hex string) (*big.Int, *big.Int) {
	bytes := HexToBytes(hex)
	halfLength := len(bytes) / 2
	aBytes := bytes[:halfLength]
	bBytes := bytes[halfLength:]

	a, b := BigIntsFromBytes(aBytes, bBytes)
	return a, b
}

// BigIntsFromBytes decodes two big.Int pointers from a byte slice.
func BigIntsFromBytes(aBytes, bBytes []byte) (*big.Int, *big.Int) {
	a, err := BytesToBigInt(aBytes)
	PanicError(err)
	b, err := BytesToBigInt(bBytes)
	PanicError(err)

	return a, b
}

// BytesToBigInt decodes a big.Int pointer from a byte slice.
func BytesToBigInt(bytes []byte) (*big.Int, error) {
	if len(bytes) != 32 {
		return nil, ErrBigIntBadBytes
	}

	bi := new(big.Int).SetBytes(bytes)
	return bi, nil
}
