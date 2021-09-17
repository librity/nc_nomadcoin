package wallet

import (
	"crypto/x509"
	"io/fs"
)

type fakeFileLayer struct {
	fakeWalletExists func() bool
}

func (f fakeFileLayer) walletExists() bool {
	return f.fakeWalletExists()
}

func (fakeFileLayer) write(name string, data []byte, perm fs.FileMode) error {
	return nil
}

func (fakeFileLayer) read(name string) ([]byte, error) {
	return x509.MarshalECPrivateKey(testWallet.privateKey)
}
