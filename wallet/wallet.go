package wallet

import (
	"crypto/ecdsa"
	"fmt"
	"os"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

var (
	w *wallet
)

const (
	walletFilepath = "./my.wallet"
)

func Start() {
	wallet := GetW()
	wallet.inspect()
}

func GetW() *wallet {
	if w == nil {
		initializeWallet()
	}

	return w
}

func initializeWallet() {
	if walletFileExists() {
		restoreWalletFromFile()
		return
	}

	createWallet()
}

func walletFileExists() bool {
	_, err := os.Stat(walletFilepath)
	walletFileMissing := os.IsNotExist(err)

	return !walletFileMissing
}

func restoreWalletFromFile() {
	key := keyFromFile(walletFilepath)
	w = restoreWallet(key)

	fmt.Println("Wallet restored from file:", walletFilepath)
}

func createWallet() {
	w = newWallet()
	keyToFile(w.privateKey, walletFilepath)

	fmt.Println("Wallet created and saved to file:", walletFilepath)
}

func newWallet() *wallet {
	privateKey := generateKey()
	publicKey := privateKey.PublicKey
	w := &wallet{privateKey, &publicKey}

	return w
}

func restoreWallet(privateKey *ecdsa.PrivateKey) *wallet {
	publicKey := privateKey.PublicKey
	w := &wallet{privateKey, &publicKey}

	return w
}

func (w *wallet) inspect() {
	fmt.Println("=== Wallet ===")
	fmt.Println("curve:", w.publicKey.Curve.Params().B)
	fmt.Println("Public key")
	fmt.Println("x:", w.publicKey.X)
	fmt.Println("y:", w.publicKey.Y)
	fmt.Println("Private key")
	fmt.Println("d:", w.privateKey.D)
	fmt.Println("hex:", keyToHex(w.privateKey))
	fmt.Println()
}
