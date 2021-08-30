package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data         string
	previousHash string
	hash         string
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
	newBlock := block{data, b.getLastHash(), ""}
	newBlock.setHash()
	b.blocks = append(b.blocks, newBlock)

}

func (b *blockchain) listBlocks() {
	for index, block := range b.blocks {
		fmt.Printf("Block number: %d\n", index+1)
		fmt.Printf("Data: %s\n", block.data)
		fmt.Printf("Previous hash: %s\n", block.previousHash)
		fmt.Printf("Hash: %s\n", block.hash)
		fmt.Println("---")
	}
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
