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

func keyToFile(key *ecdsa.PrivateKey, filePath string) {
	keyBytes := keyToBytes(key)
	secureFilePerm := fs.FileMode(0600)

	err := os.WriteFile(filePath, keyBytes, secureFilePerm)
	utils.HandleError(err)
}

func keyFromFile(filePath string) (key *ecdsa.PrivateKey) {
	keyBytes, err := os.ReadFile(filePath)
	utils.HandleError(err)

	key = keyFromBytes(keyBytes)
	return
}

func keyToBytes(key *ecdsa.PrivateKey) []byte {
	keyBytes, err := x509.MarshalECPrivateKey(key)
	utils.HandleError(err)

	return keyBytes
}

func keyToHex(key *ecdsa.PrivateKey) string {
	bytes := keyToBytes(key)

	hex := fmt.Sprintf("%x", bytes)
	return hex
}

func keyFromBytes(keyBytes []byte) *ecdsa.PrivateKey {
	key, err := x509.ParseECPrivateKey(keyBytes)
	utils.HandleError(err)

	return key
}

func keyFromHex(keyHex string) *ecdsa.PrivateKey {
	keyBytes := utils.HexToBytes(keyHex)
	key := keyFromBytes(keyBytes)

	return key
}
