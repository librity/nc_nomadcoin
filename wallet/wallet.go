package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"fmt"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
	Address    string
}

var (
	w     *wallet
	curve = elliptic.P256()
)

const (
	walletFilepath = "./my.wallet"
)

func GetW() *wallet {
	if w == nil {
		initializeWallet()
	}

	return w
}

func GetAddress() string {
	return GetW().Address
}

func initializeWallet() {
	if files.walletExists() {
		initializeWalletFromFile()
		return
	}

	createWallet()
}

func initializeWalletFromFile() {
	key := keyFromFile(walletFilepath)
	w = newWallet(key)

	fmt.Println("👛 Wallet initialized from file:", walletFilepath)
	w.prettyAddress()
}

func createWallet() {
	key := generateKey()
	w = newWallet(key)
	keyToFile(w.privateKey, walletFilepath)

	fmt.Println("👛 Wallet created and saved to file:", walletFilepath)
	w.prettyAddress()
}

func newWallet(key *ecdsa.PrivateKey) *wallet {
	w := &wallet{
		privateKey: key,
		PublicKey:  &key.PublicKey,
		Address:    addressFromKey(key),
	}

	return w
}

func (w *wallet) Inspect() {
	fmt.Println("=== Wallet ===")
	fmt.Println("curve:", w.PublicKey.Curve.Params().B)
	fmt.Println("address:", w.Address)

	fmt.Println("public_key:")
	fmt.Println("	x:", w.PublicKey.X)
	fmt.Println("	y:", w.PublicKey.Y)

	fmt.Println("private_key:")
	fmt.Println("	d:", w.privateKey.D)
	fmt.Println()
}

func (w *wallet) prettyAddress() {
	fmt.Println("👛 Wallet address:", w.Address)
}
