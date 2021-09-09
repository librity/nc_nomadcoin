package wallet

import (
	"fmt"
	"math/big"

	"github.com/librity/nc_nomadcoin/utils"
)

type Signature struct {
	R *big.Int
	S *big.Int
}

func sinatureToHex(r *big.Int, s *big.Int) string {
	bytes := append(r.Bytes(), s.Bytes()...)
	hex := fmt.Sprintf("%x", bytes)

	return hex
}

func hexToSignature(signatureHex string) *Signature {
	signatureBytes := utils.HexToBytes(signatureHex)
	halfLength := len(signatureBytes) / 2
	rBytes := signatureBytes[:halfLength]
	sBytes := signatureBytes[halfLength:]

	r, err := utils.BytesToBigInt(rBytes)
	utils.HandleError(err)
	s, err := utils.BytesToBigInt(sBytes)
	utils.HandleError(err)

	signature := &Signature{r, s}
	return signature
}
