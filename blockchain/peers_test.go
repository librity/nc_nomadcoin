package blockchain

import (
	"testing"

	"github.com/librity/nc_nomadcoin/utils"
)

func TestAddPeerBlock(t *testing.T) {
	peerBlock := Block{
		Hash:         "3",
		PreviousHash: "2",
		Difficulty:   4,
		Height:       3,
	}

	t.Run("Should append peer block to blockchain", func(t *testing.T) {
		chain := &blockchain{
			LastHash:   "2",
			Difficulty: 3,
			Height:     2,
		}
		chain.addPeerBlock(&peerBlock)

		utils.FailIfDifferent(t, peerBlock.Hash, chain.LastHash)
		utils.FailIfDifferent(t, peerBlock.Difficulty, chain.Difficulty)
		utils.FailIfDifferent(t, peerBlock.Height, chain.Height)

	})

	t.Run("", func(t *testing.T) {

	})

}

func Test(t *testing.T) {

	t.Run("", func(t *testing.T) {

	})

}
