package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"

	"github.com/librity/nc_nomadcoin/utils"
)

type signature struct {
	R *big.Int
	S *big.Int
}

func HexSign(hash string) string {
	payloadBytes := utils.HexToBytes(hash)
	privateKey := GetW().privateKey

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, payloadBytes)
	utils.HandleError(err)

	signHex := utils.BigIntsToHex(r, s)
	return signHex
}

func Verify(hash, signHex, address string) bool {
	signature := signFromHex(signHex)
	publicKey := addressToPublicKey(address)
	hashBytes := utils.HexToBytes(hash)

	isValid := ecdsa.Verify(publicKey, hashBytes, signature.R, signature.S)
	return isValid
}

func signFromHex(hex string) *signature {
	r, s := utils.BigIntsFromHex(hex)

	signature := newSignature(r, s)
	return signature
}

func newSignature(r, s *big.Int) *signature {
	signature := &signature{
		R: r,
		S: s,
	}

	return signature
}
