package blockchain

import (
	"errors"
	"fmt"

	"github.com/librity/nc_nomadcoin/utils"
)

var (
	ErrBlockNotFound = errors.New("block not found")
)

func GetBlocks() []*Block {
	return getBlocks(getBC())
}

func getBlocks(chain *blockchain) []*Block {
	var blocks []*Block
	currentHash := chain.LastHash

	for {
		block, err := FindBlock(currentHash)
		utils.PanicError(err)

		blocks = append(blocks, block)
		currentHash = block.PreviousHash

		if currentHash == "" {
			break
		}
	}

	return blocks
}

func FindBlock(hash string) (*Block, error) {
	rawBlock := storage.LoadBlock(hash)
	if rawBlock == nil {
		return nil, ErrBlockNotFound
	}

	block := blockFromBytes(rawBlock)
	return block, nil
}

func GetLastNBlocks(n int) []*Block {
	return getLastNBlocks(getBC(), n)
}

func getLastNBlocks(chain *blockchain, n int) []*Block {
	var blocks []*Block
	currentHash := chain.LastHash

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

func GetLastBlock() *Block {
	return getLastBlock(getBC())
}

func getLastBlock(chain *blockchain) *Block {
	lastHash := getBC().LastHash
	lastBlock, err := FindBlock(lastHash)
	utils.PanicError(err)

	return lastBlock
}

func InspectBlocks() {
	fmt.Println("=== Blocks ===")

	blocks := GetBlocks()
	for _, block := range blocks {
		block.inspect()
	}

	fmt.Println("")
}
