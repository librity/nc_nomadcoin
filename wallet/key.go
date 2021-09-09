package wallet

import (
	"crypto/ecdsa"
	"crypto/x509"
	"fmt"

	"github.com/librity/nc_nomadcoin/utils"
)

func keyToHex(privateKey *ecdsa.PrivateKey) string {
	bytes, err := x509.MarshalECPrivateKey(privateKey)
	utils.HandleError(err)

	hex := fmt.Sprintf("%x", bytes)
	return hex
}

func hexToKey(keyHex string) *ecdsa.PrivateKey {
	keyBytes := utils.HexToBytes(keyHex)
	privateKey, err := x509.ParseECPrivateKey(keyBytes)
	utils.HandleError(err)

	return privateKey
}
