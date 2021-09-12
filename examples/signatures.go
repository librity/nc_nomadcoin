package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/librity/nc_nomadcoin/utils"
)

type Message struct {
	From    string
	To      string
	Content string
}

type Signature struct {
	R *big.Int
	S *big.Int
}

func main() {
	hash := buildMessage()
	privateKey, publicKey := generateKeys()
	// tamperPrivateKey(privateKey)
	// tamperPublicKey(publicKey)
	signature := sign(privateKey, hash)
	// tamperSignature(signature)
	// tamperHash(hash)
	verify(publicKey, hash, signature)
}

func buildMessage() []byte {
	message := Message{"lior", "waiFu", "kpop makes me nauseous"}
	hexHash := utils.HexHash(message)
	hash := utils.HexToBytes(hexHash)

	fmt.Println("=== Message ===")
	fmt.Println("Struct:", message)
	fmt.Println("Hash:", hexHash)
	fmt.Println("Bytes:", hash)

	return hash
}

func generateKeys() (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.PanicError(err)
	publicKey := &privateKey.PublicKey

	fmt.Println("=== Public key ===")
	fmt.Println("x:", publicKey.X)
	fmt.Println("y:", publicKey.Y)
	fmt.Println("=== Private key ===")
	fmt.Println("d:", privateKey.D)

	return privateKey, publicKey
}

func sign(privateKey *ecdsa.PrivateKey, hash []byte) *Signature {
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash)
	utils.PanicError(err)
	signature := &Signature{r, s}

	fmt.Println("=== Signature ===")
	fmt.Println("R:", signature.R)
	fmt.Println("S:", signature.S)

	return signature
}

func verify(publicKey *ecdsa.PublicKey, hash []byte, signature *Signature) {
	isValid := ecdsa.Verify(publicKey, hash, signature.R, signature.S)

	fmt.Println("=== Verification ===")
	if isValid {
		fmt.Println("🔐 Valid signature! 🛡️")
		return
	}

	fmt.Println("🔓 Invalid signature: Authorities alerted! 👨‍💻⚠️")
}

func tamperPrivateKey(key *ecdsa.PrivateKey) {
	fmt.Println("=== Tampering Private Key ===")
	fmt.Println("Before:", key.D)

	tamper := big.NewInt(1)
	key.D = tamper.Add(key.D, tamper)
	fmt.Println("After:", key.D)
}

func tamperPublicKey(key *ecdsa.PublicKey) {
	fmt.Println("=== Tampering Public Key ===")
	fmt.Println("Before:", key.X, key.Y)

	tamper := big.NewInt(1)
	key.X = tamper.Add(key.X, tamper)
	tamper = big.NewInt(-1)
	key.Y = tamper.Add(key.Y, tamper)
	fmt.Println("After:", key.X, key.Y)
}

func tamperSignature(signature *Signature) {
	fmt.Println("=== Tampering Signature ===")
	fmt.Println("Before:", signature.R, signature.S)

	tamper := big.NewInt(1)
	signature.R = tamper.Add(signature.R, tamper)
	tamper = big.NewInt(-1)
	signature.S = tamper.Add(signature.S, tamper)
	fmt.Println("After:", signature.R, signature.S)
}

func tamperHash(hash []byte) {
	fmt.Println("=== Tampering Hash ===")
	fmt.Println("Before:", hash)

	hash['W'%31] = byte(9)
	hash['T'%31] = byte(1)
	hash['C'%31] = byte(1)
	fmt.Println("After:", hash)
}
