package wallet

import (
	"math/big"

	"github.com/librity/nc_nomadcoin/utils"
)

type Signature struct {
	R *big.Int
	S *big.Int
}

func signFromHex(hex string) *Signature {
	bytes := utils.HexToBytes(hex)
	halfLength := len(bytes) / 2
	rBytes := bytes[:halfLength]
	sBytes := bytes[halfLength:]

	signature := signFromBytes(rBytes, sBytes)
	return signature
}

func signFromBytes(rBytes []byte, sBytes []byte) *Signature {
	r, err := utils.BytesToBigInt(rBytes)
	utils.HandleError(err)
	s, err := utils.BytesToBigInt(sBytes)
	utils.HandleError(err)

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
