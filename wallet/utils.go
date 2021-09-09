package wallet

import (
	"fmt"
	"math/big"
)

func bigIntsToHex(a *big.Int, b *big.Int) string {
	bytes := bigIntsToBytes(a, b)
	hex := fmt.Sprintf("%x", bytes)

	return hex
}

func bigIntsToBytes(a *big.Int, b *big.Int) []byte {
	bytes := append(a.Bytes(), b.Bytes()...)

	return bytes
}
