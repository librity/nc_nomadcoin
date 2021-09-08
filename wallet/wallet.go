package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"

	"github.com/librity/nc_nomadcoin/utils"
)

func Start() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleError(err)
	publicKey := privateKey.PublicKey

	fmt.Println("=== Public key ===")
	fmt.Println("x:", publicKey.X)
	fmt.Println("y:", publicKey.Y)
	fmt.Println("=== Private key ===")
	fmt.Println(privateKey.D)
}
