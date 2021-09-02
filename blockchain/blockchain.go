package blockchain

import (
	"errors"
	"fmt"
	"sync"
)

var (
	ErrNotFound = errors.New("block not found")
)

type blockchain struct {
	blocks []*Block
}

var (
	b    *blockchain
	once sync.Once
)

func Get() *blockchain {
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

func (b *blockchain) AllBlocks() []*Block {
	return b.blocks
}

func (b *blockchain) GetBlock(height int) (*Block, error) {
	if height > len(b.blocks) {
		return nil, ErrNotFound
	}
	if height < 1 {
		return nil, ErrNotFound
	}

	return b.blocks[height-1], nil
}

func initializeBlockchain() {
	b = &blockchain{}
	b.AddBlock("The Times 03/Jan/2009 Chancellor on brink of second bailout for banks")
}

func createBlock(data string) *Block {
	newBlock := Block{getBlockHeight(), data, getLastHash(), ""}
	newBlock.setHash()
	return &newBlock
}

func getLastHash() string {
	b := Get()
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
