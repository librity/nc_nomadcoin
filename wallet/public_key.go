package wallet

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/librity/nc_nomadcoin/utils"
)

func addressFromKey(key *ecdsa.PrivateKey) string {
	publicKey := &key.PublicKey
	address := addressFromPublicKey(publicKey)

	return address
}

func addressFromPublicKey(publicKey *ecdsa.PublicKey) string {
	address := utils.BigIntsToHex(publicKey.X, publicKey.Y)

	return address
}

func addressToPublicKey(address string) *ecdsa.PublicKey {
	x, y := utils.BigIntsFromHex(address)

	publicKey := newPublicKey(x, y)
	return publicKey
}

func newPublicKey(x, y *big.Int) *ecdsa.PublicKey {
	publicKey := &ecdsa.PublicKey{
		Curve: curve,
		X:     x,
		Y:     y,
	}

	return publicKey
}
