package blockchain

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Number       int    `json:"number"`
	Data         string `json:"data"`
	PreviousHash string `json:"previousHash,omitempty"`
	Hash         string `json:"hash"`
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

// Stringer interface: https://pkg.go.dev/fmt#Stringer
func (b Block) String() string {
	s := fmt.Sprintln("Block number:", fmt.Sprint(b.Number)) +
		fmt.Sprintln("Data:", b.Data)
	if b.PreviousHash != "" {
		s = s + fmt.Sprintln("Previous hash:", b.PreviousHash)
	}
	s = s + fmt.Sprintln("Hash:", b.Hash) +
		fmt.Sprintln("---")

	return s
}

func (b *Block) listBlock() {
	fmt.Printf("Block number: %d\n", b.Number)
	fmt.Printf("Data: %s\n", b.Data)
	if b.PreviousHash != "" {
		fmt.Printf("Previous hash: %s\n", b.PreviousHash)
	}
	fmt.Printf("Hash: %s\n", b.Hash)
	fmt.Println("---")
}
