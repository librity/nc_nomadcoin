package blockchain

import (
	"sync"
	"testing"

	"github.com/librity/nc_nomadcoin/utils"
)

func makeFakeBC() *blockchain {
	lastHash := utils.RandomHash()

	return newBC(lastHash, 1, 1)
}

func TestGetBC(t *testing.T) {

	t.Run("Should create blockchain when it doesn't exist", func(t *testing.T) {
		onceBC = *new(sync.Once)
		storage = fakeStorageLayer{
			fakeLoadChain: func() []byte { return nil },
		}

		result := getBC()
		if result != bc {
			t.Errorf("Expected %v, got %v", bc, result)
		}

		if result.Height != 1 {
			t.Errorf("Expected %v, got %v", 1, result.Height)
		}
	})

	t.Run("Should restore blockchain when it does exist", func(t *testing.T) {
		onceBC = *new(sync.Once)
		lastHash := utils.RandomHash()
		storage = fakeStorageLayer{
			fakeLoadChain: func() []byte {
				chain := &blockchain{
					LastHash:  lastHash,
					Height:    2,
					Dificulty: 1,
				}
				return utils.ToGob(chain)
			},
		}

		result := getBC()
		if result.LastHash != lastHash {
			t.Errorf("Expected %v, got %v", lastHash, result.LastHash)
		}

		if result.Height != 2 {
			t.Errorf("Expected %v, got %v", 2, result.Height)
		}
	})

}

func Test(t *testing.T) {

	t.Run("", func(t *testing.T) {

	})

}
