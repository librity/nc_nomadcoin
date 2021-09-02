package blockchain

import (
	"errors"
	"sync"
)

var (
	ErrNotFound = errors.New("block not found")
)

type blockchain struct {
	LastHash string `json:"newestHash"`
	Height   int    `json:"height"`
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
	block := createBlock(data, b.LastHash, b.Height)

	b.LastHash = block.Hash
	b.Height = block.Height
}

func (b *blockchain) ListBlocks() {
}

func (b *blockchain) AllBlocks() {
}

func (b *blockchain) GetBlock(height int) {

}

func initializeBlockchain() {
	b = &blockchain{"", 0}
	b.AddBlock("The Times 03/Jan/2009 Chancellor on brink of second bailout for banks")
}
