package blockchain

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Data         string
	PreviousHash string
	Hash         string
}

func (b *Block) setHash() {
	b.Hash = b.generateHash()
}

func (b *Block) generateHash() string {
	preHash := b.Data + b.PreviousHash
	preHashBytes := []byte(preHash)
	rawHash := sha256.Sum256(preHashBytes)
	hexHash := fmt.Sprintf("%x", rawHash)

	return hexHash
}

func (b *Block) listBlock() {
	fmt.Printf("Data: %s\n", b.Data)
	fmt.Printf("Previous hash: %s\n", b.PreviousHash)
	fmt.Printf("Hash: %s\n", b.Hash)
}
