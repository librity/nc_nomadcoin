package blockchain

import (
	"sync"
	"testing"

	"github.com/librity/nc_nomadcoin/utils"
)

func makeFakeBC() *blockchain {
	chain := &blockchain{
		LastHash:   utils.RandomHash(),
		Height:     1,
		Difficulty: 1,
	}

	return chain
}

func TestGetBC(t *testing.T) {

	t.Run("Should create blockchain when it doesn't exist", func(t *testing.T) {
		onceBC = *new(sync.Once)
		storage = fakeStorageLayer{
			fakeLoadChain: func() []byte { return nil },
		}

		result := getBC()
		utils.FailIfDifferent(t, bc, result)
		utils.FailIfDifferent(t, 1, result.Height)

	})

	t.Run("Should restore blockchain when it does exist", func(t *testing.T) {
		onceBC = *new(sync.Once)
		lastHash := utils.RandomHash()
		storage = fakeStorageLayer{
			fakeLoadChain: func() []byte {
				chain := &blockchain{
					LastHash:   lastHash,
					Height:     2,
					Difficulty: 1,
				}
				return utils.ToGob(chain)
			},
		}

		result := getBC()
		utils.FailIfDifferent(t, lastHash, result.LastHash)
		utils.FailIfDifferent(t, 2, result.Height)

	})

}
