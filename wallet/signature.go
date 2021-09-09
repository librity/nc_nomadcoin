package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"

	"github.com/librity/nc_nomadcoin/utils"
)

type Signature struct {
	R *big.Int
	S *big.Int
}

func hexSign(hash string, w *wallet) string {
	payloadBytes := utils.HexToBytes(hash)

	r, s, err := ecdsa.Sign(rand.Reader, w.privateKey, payloadBytes)
	utils.HandleError(err)

	signHex := utils.BigIntsToHex(r, s)
	return signHex
}

func verify(signHex, hash, address string) bool {
	signature := signFromHex(signHex)
	publicKey := addressToPublicKey(address)
	hashBytes := utils.HexToBytes(hash)

	isValidSignature := ecdsa.Verify(publicKey, hashBytes, signature.R, signature.S)
	return isValidSignature
}

func signFromHex(hex string) *Signature {
	r, s := utils.BigIntsFromHex(hex)

	signature := newSignature(r, s)
	return signature
}

func newSignature(r *big.Int, s *big.Int) *Signature {
	signature := &Signature{
		R: r,
		S: s,
	}

	return signature
}
