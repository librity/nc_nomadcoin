package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"fmt"
	"io/fs"
	"os"

	"github.com/librity/nc_nomadcoin/utils"
)

func generateKey() *ecdsa.PrivateKey {
	curve := elliptic.P256()
	key, err := ecdsa.GenerateKey(curve, rand.Reader)
	utils.HandleError(err)

	return key
}

func keyFromFile(filePath string) *ecdsa.PrivateKey {
	fileBytes, err := os.ReadFile(filePath)
	utils.HandleError(err)

	keyHex := utils.BytesToHex(fileBytes)
	key := hexToKey(keyHex)

	return key
}

func keyToFile(key *ecdsa.PrivateKey, filePath string) {
	keyHex := keyToHex(key)
	keyBytes := utils.HexToBytes(keyHex)
	secureFilePerm := fs.FileMode(0600)

	os.WriteFile(filePath, keyBytes, secureFilePerm)
}

func keyToHex(key *ecdsa.PrivateKey) string {
	bytes, err := x509.MarshalECPrivateKey(key)
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
