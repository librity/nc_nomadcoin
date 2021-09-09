package wallet

import (
	"crypto/ecdsa"
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
	walletFile = "my.wallet"
)

func GetW() *wallet {
	if w == nil {
		initializeWallet()
	}

	return w
}

func initializeWallet() {
	if hasWalletFile() {
		restoreFromFile()
		return
	}

	createWallet()
}

func hasWalletFile() bool {
	_, err := os.Stat(walletFile)
	walletFileMissing := os.IsNotExist(err)

	return !walletFileMissing
}

func restoreFromFile() {

}

func createWallet() {

}
