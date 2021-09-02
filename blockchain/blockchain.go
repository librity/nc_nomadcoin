package blockchain

import (
	"fmt"
	"sync"
)

type blockchain struct {
	blocks []*Block
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
	for _, block := range b.blocks {
		// block.listBlock()
		fmt.Print(block)
	}
}

func (b *blockchain) GetAllBlocks() []*Block {
	return b.blocks
}

func (b *blockchain) GetBlock(height int) *Block {
	return b.blocks[height-1]
}

func initializeBlockchain() {
	b = &blockchain{}
	b.AddBlock("Genesis block.")
}

func createBlock(data string) *Block {
	newBlock := Block{getBlockHeight(), data, getLastHash(), ""}
	newBlock.setHash()
	return &newBlock
}

func getLastHash() string {
	b := GetBlockchain()
	if b.isFirstBlock() {
		return ""
	}

	lastBlockIndex := len(b.blocks) - 1
	lastHash := b.blocks[lastBlockIndex].Hash
	return lastHash
}

func getBlockHeight() int {
	return len(b.blocks) + 1
}

func (b *blockchain) isFirstBlock() bool {
	return len(b.blocks) == 0
}
