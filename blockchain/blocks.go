package blockchain

import (
	"errors"
	"fmt"

	"github.com/librity/nc_nomadcoin/db"
	"github.com/librity/nc_nomadcoin/utils"
)

var (
	ErrBlockNotFound = errors.New("block not found")
)

func GetBlocks() []*Block {
	var blocks []*Block
	currentHash := getBC().lastHash

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

func FindBlock(hash string) (*Block, error) {
	rawBlock := db.LoadBlock(hash)
	if rawBlock == nil {
		return nil, ErrBlockNotFound
	}

	block := blockFromBytes(rawBlock)
	return block, nil
}

func GetLastNBlocks(n int) []*Block {
	var blocks []*Block
	currentHash := getBC().lastHash

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
	lastHash := getBC().lastHash
	lastBlock, err := FindBlock(lastHash)
	utils.PanicError(err)

	return lastBlock
}

func Replace(blocks []*Block) {
	chain := getBC()
	resetBC()

	for _, block := range blocks {
		block.save()
	}

	lastBlock := blocks[0]
	chain.reference(lastBlock)

	fmt.Println("⛓️  Blockchain replaced succesfully.")
}

func InspectBlocks() {
	fmt.Println("=== Blocks ===")

	blocks := GetBlocks()
	for _, block := range blocks {
		block.inspect()
	}

	fmt.Println("")
}
