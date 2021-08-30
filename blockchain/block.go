package blockchain

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data         string
	previousHash string
	hash         string
}

func (b *block) setHash() {
	b.hash = b.generateHash()
}

func (b *block) generateHash() string {
	preHash := b.data + b.previousHash
	preHashBytes := []byte(preHash)
	rawHash := sha256.Sum256(preHashBytes)
	hexHash := fmt.Sprintf("%x", rawHash)

	return hexHash
}

func (b *block) listBlock() {
	fmt.Printf("Data: %s\n", b.data)
	fmt.Printf("Previous hash: %s\n", b.previousHash)
	fmt.Printf("Hash: %s\n", b.hash)
}
