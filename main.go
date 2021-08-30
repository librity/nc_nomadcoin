package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data         string
	hash         string
	previousHash string
}

type blockchain struct {
	blocks []block
}

func main() {
	welcome()
	demo()
}

func welcome() {
	fmt.Println("Welcome to Nomad Coin!")
	fmt.Println("---")
}

func demo() {
	chain := blockchain{}
	chain.addBlock("Genesis block.")
	chain.addBlock("Second block.")
	chain.addBlock("Third block.")
	chain.listBlocks()
}

func (b *blockchain) isFirstBlock() bool {
	return len(b.blocks) == 0
}

func (b *blockchain) getLastHash() string {
	if b.isFirstBlock() {
		return ""
	}

	lastBlockIndex := len(b.blocks) - 1
	lastHash := b.blocks[lastBlockIndex].hash
	return lastHash
}

func (b *blockchain) addBlock(data string) {
	newBlock := block{data, "", b.getLastHash()}
	newBlock.setHash()
	// Find previous hash

}

func (b *blockchain) listBlocks() {

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
