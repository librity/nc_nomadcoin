package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/x509"
	"fmt"
	"io/fs"
	"os"

	"github.com/librity/nc_nomadcoin/utils"
)

func generateKey() *ecdsa.PrivateKey {
	key, err := ecdsa.GenerateKey(curve, rand.Reader)
	utils.PanicError(err)

	return key
}

func keyToFile(key *ecdsa.PrivateKey, filePath string) {
	keyBytes := keyToBytes(key)
	secureFilePerm := fs.FileMode(0600)

	err := os.WriteFile(filePath, keyBytes, secureFilePerm)
	utils.PanicError(err)
}

func keyFromFile(filePath string) (key *ecdsa.PrivateKey) {
	keyBytes, err := os.ReadFile(filePath)
	utils.PanicError(err)

	key = keyFromBytes(keyBytes)
	return
}

func keyToBytes(key *ecdsa.PrivateKey) []byte {
	keyBytes, err := x509.MarshalECPrivateKey(key)
	utils.PanicError(err)

	return keyBytes
}

func keyToHex(key *ecdsa.PrivateKey) string {
	bytes := keyToBytes(key)

	hex := fmt.Sprintf("%x", bytes)
	return hex
}

func keyFromBytes(keyBytes []byte) *ecdsa.PrivateKey {
	key, err := x509.ParseECPrivateKey(keyBytes)
	utils.PanicError(err)

	return key
}

func keyFromHex(keyHex string) *ecdsa.PrivateKey {
	keyBytes := utils.HexToBytes(keyHex)
	key := keyFromBytes(keyBytes)

	return key
}
