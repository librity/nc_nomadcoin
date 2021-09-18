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
		if blocksType != expectedType {
			t.Errorf("Expected %v, got %v", expectedType, blocksType)
		}
	})

	t.Run("Should return all blocks in sequence", func(t *testing.T) {
		blockControl := 0
		storage = fakeStorageLayer{
			fakeLoadBlock: func() []byte {
				fakeBlocks := []*Block{
					{Hash: "3", PreviousHash: "2"},
					{Hash: "2", PreviousHash: "1"},
					{Hash: "1", PreviousHash: ""},
				}

				chosenBlock := fakeBlocks[blockControl]
				blockControl++
				return utils.ToGob(chosenBlock)
			},
		}
		chain := &blockchain{LastHash: "3"}

		blocks := getBlocks(chain)
		blocksLen := len(blocks)
		if blocksLen != 3 {
			t.Errorf("Expected %v, got %v", 3, blocksLen)
		}

		if blocks[0].Hash != "3" {
			t.Errorf("Expected %v, got %v", "3", blocks[0].Hash)
		}
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
		if err != nil {
			t.Errorf("Expected %v, got %v", nil, err)
		}

		if block.Hash != fakeBlock.Hash {
			t.Errorf("Expected %v, got %v", fakeBlock.Hash, block.Hash)
		}
	})

	t.Run("Should return ErrBlockNotFound when block doesn't exist", func(t *testing.T) {
		storage = fakeStorageLayer{
			fakeLoadBlock: func() []byte { return nil },
		}

		_, err := FindBlock("test")
		if err != ErrBlockNotFound {
			t.Errorf("Expected %v, got %v", ErrBlockNotFound, err)
		}
	})

}
