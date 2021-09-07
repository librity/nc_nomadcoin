package blockchain

import (
	"fmt"
)

func Blocks() []*Block {
	var blocks []*Block
	currentHash := Get().LastHash

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

func LastNBlocks(n int) []*Block {
	var blocks []*Block
	currentHash := Get().LastHash

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

func InspectBlocks() {
	fmt.Println("=== Blocks ===")

	blocks := Blocks()
	for _, block := range blocks {
		block.inspect()
	}

	fmt.Println("")
}
