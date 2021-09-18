package blockchain

import (
	"reflect"
	"testing"

	"github.com/librity/nc_nomadcoin/utils"
)

func TestGetBlocks(t *testing.T) {

	t.Run("Should return a slice of *Blocks", func(t *testing.T) {
		fakeBlock := makeFakeBlock()
		storage = fakeStorageLayer{
			fakeLoadBlock: func() []byte { return utils.ToGob(fakeBlock) },
		}
		chain := &blockchain{LastHash: fakeBlock.Hash}

		blocks := getBlocks(chain)
		blocksType := reflect.TypeOf(blocks)
		expectedType := reflect.TypeOf([]*Block{})
		utils.FailIfDifferent(t, expectedType, blocksType)

	})

	t.Run("Should return all blocks in sequence", func(t *testing.T) {
		blockControl := 0
		fakeBlocks := []*Block{
			{Hash: "3", PreviousHash: "2"},
			{Hash: "2", PreviousHash: "1"},
			{Hash: "1", PreviousHash: ""},
		}
		storage = fakeStorageLayer{
			fakeLoadBlock: func() []byte {
				defer func() { blockControl++ }()

				return utils.ToGob(fakeBlocks[blockControl])
			},
		}
		chain := &blockchain{LastHash: "3"}

		blocks := getBlocks(chain)
		blocksLen := len(blocks)
		utils.FailIfDifferent(t, 3, blocksLen)
		utils.FailIfDifferent(t, "3", blocks[0].Hash)

	})

}

func TestFindBlock(t *testing.T) {
	fakeBlock := makeFakeBlock()
	fakeBlockBytes := utils.ToGob(fakeBlock)

	t.Run("Should return fakeBlock when it exists", func(t *testing.T) {
		storage = fakeStorageLayer{
			fakeLoadBlock: func() []byte { return fakeBlockBytes },
		}

		block, err := FindBlock(fakeBlock.Hash)
		utils.FailIfDifferent(t, nil, err)
		utils.FailIfDifferent(t, fakeBlock.Hash, block.Hash)

	})

	t.Run("Should return ErrBlockNotFound when block doesn't exist", func(t *testing.T) {
		storage = fakeStorageLayer{
			fakeLoadBlock: func() []byte { return nil },
		}

		_, err := FindBlock("test")
		utils.FailIfDifferent(t, ErrBlockNotFound, err)

	})

}
