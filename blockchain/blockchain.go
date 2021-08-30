package blockchain

import "fmt"

type blockchain struct {
	blocks []block
}

var b *blockchain

func GetBlockchain() *blockchain {
	if b == nil {
		b = &blockchain{}
	}
	return b
}

func (b *blockchain) AddBlock(data string) {
	newBlock := block{data, b.getLastHash(), ""}
	newBlock.setHash()
	b.blocks = append(b.blocks, newBlock)
}

func (b *blockchain) ListBlocks() {
	for index, block := range b.blocks {
		fmt.Printf("Block number: %d\n", index+1)
		block.listBlock()
		fmt.Println("---")
	}
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
