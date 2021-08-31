package blockchain

import (
	"fmt"
	"sync"
)

type blockchain struct {
	blocks []*block
}

var b *blockchain
var once sync.Once

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(initializeBlockchain)
	}
	return b
}

func (b *blockchain) AddBlock(data string) {
	newBlock := createBlock(data)
	b.blocks = append(b.blocks, newBlock)
}

func (b *blockchain) ListBlocks() {
	for index, block := range b.blocks {
		fmt.Printf("Block number: %d\n", index+1)
		block.listBlock()
		fmt.Println("---")
	}
}

func initializeBlockchain() {
	b = &blockchain{}
	b.AddBlock("Genesis block.")
}

func createBlock(data string) *block {
	newBlock := block{data, getLastHash(), ""}
	newBlock.setHash()
	return &newBlock
}

func getLastHash() string {
	b := GetBlockchain()
	if b.isFirstBlock() {
		return ""
	}

	lastBlockIndex := len(b.blocks) - 1
	lastHash := b.blocks[lastBlockIndex].hash
	return lastHash
}

func (b *blockchain) isFirstBlock() bool {
	return len(b.blocks) == 0
}
