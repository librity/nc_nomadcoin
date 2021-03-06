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

func HexSign(hash string, w *wallet) string {
	payloadBytes := utils.HexToBytes(hash)

	r, s, err := ecdsa.Sign(rand.Reader, w.privateKey, payloadBytes)
	utils.PanicError(err)

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
