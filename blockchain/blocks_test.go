package blockchain

import (
	"testing"

	"github.com/librity/nc_nomadcoin/utils"
)

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
