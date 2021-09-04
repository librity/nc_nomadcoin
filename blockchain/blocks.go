package blockchain

import (
	"fmt"
)

func (b *blockchain) Blocks() []*Block {
	var blocks []*Block
	currentHash := b.LastHash

	for {
		block, _ := FindBlock(currentHash)
		blocks = append(blocks, block)
		currentHash = block.PreviousHash

		if currentHash == "" {
			break
		}
	}

	return blocks
}

func (b *blockchain) LastNBlocks(n int) []*Block {
	var blocks []*Block
	currentHash := b.LastHash
	for i := 0; i < n; i++ {
		block, _ := FindBlock(currentHash)
		blocks = append(blocks, block)
		currentHash = block.PreviousHash

		if currentHash == "" {
			break
		}
	}

	return blocks
}

func (b *blockchain) ListBlocks() {
	fmt.Println("=== Blocks ===")

	blocks := b.Blocks()
	for _, block := range blocks {
		block.inspect()
	}

	fmt.Println("")
}
