package wallet

import (
	"io/fs"
	"os"
)

type fileLayerI interface {
	walletExists() bool
	write(name string, data []byte, perm fs.FileMode) error
	read(name string) ([]byte, error)
}

type fileLayerS struct{}

var (
	files fileLayerI = fileLayerS{}
)

func (fileLayerS) walletExists() bool {
	_, err := os.Stat(walletFilepath)
	walletFileMissing := os.IsNotExist(err)

	return !walletFileMissing
}

func (fileLayerS) write(name string, data []byte, perm fs.FileMode) error {
	err := os.WriteFile(name, data, perm)

	return err
}

func (fileLayerS) read(name string) ([]byte, error) {
	keyBytes, err := os.ReadFile(name)

	return keyBytes, err
}
