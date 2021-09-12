package main

import (
	"crypto/ecdsa"
	"crypto/x509"
	"fmt"
	"math/big"

	"github.com/librity/nc_nomadcoin/utils"
)

type Signature struct {
	R *big.Int
	S *big.Int
}

const (
	keyHex       = "3077020101042055958ca595187eca2fb0ccf24be179d9fa03aabf10a4145dfb1cbb58acf7c0f3a00a06082a8648ce3d030107a14403420004d6c4bab87c24e78b4753f6189d304ed733587f4a4eb7222fbbe21652e4e1cb8cee2053f20d39acc07797e3fb84ab80e2a1b6ec0220246dd593f440c1f2a956b9"
	signatureHex = "b029962795a3ab359a879772e77188d3831fc9e125864e7c88bdd798647a9289eb3385ff6c6e53992e64e0ac810a659b966b2cf1b070301a89a2416ef9518a5f"
	messageHex   = "be772bf8f967ddcfea9d7c19020d7a3d6759bb9ddb7ec79cc5b014a499556c3a"
)

func main() {
	privateKey := hexToKey(keyHex)
	publicKey := &privateKey.PublicKey
	signature := hexToSignature(signatureHex)
	messageHash := utils.HexToBytes(messageHex)

	verify(publicKey, messageHash, signature)
}

func verify(publicKey *ecdsa.PublicKey, messageHash []byte, signature *Signature) {
	isValid := ecdsa.Verify(publicKey, messageHash, signature.R, signature.S)

	fmt.Println("=== Verification ===")
	if isValid {
		fmt.Println("🔐 Valid signature! 🛡️")
		return
	}

	fmt.Println("🔓 Invalid signature: Authorities alerted! 👨‍💻⚠️")
}

func keyToHex(privateKey *ecdsa.PrivateKey) string {
	bytes, err := x509.MarshalECPrivateKey(privateKey)
	utils.PanicError(err)

	hex := fmt.Sprintf("%x", bytes)
	return hex
}

func hexToKey(keyHex string) *ecdsa.PrivateKey {
	keyBytes := utils.HexToBytes(keyHex)
	privateKey, err := x509.ParseECPrivateKey(keyBytes)
	utils.PanicError(err)

	return privateKey
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
	utils.PanicError(err)
	s, err := utils.BytesToBigInt(sBytes)
	utils.PanicError(err)

	signature := &Signature{r, s}
	return signature
}
