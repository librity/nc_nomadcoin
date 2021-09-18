package blockchain

import (
	"sync"
	"testing"

	"github.com/librity/nc_nomadcoin/utils"
)

func TestAddPeerBlock(t *testing.T) {
	tx := makeFakeTx("test")
	peerBlock := Block{
		Hash:         "3",
		PreviousHash: "2",
		Difficulty:   4,
		Height:       3,
		Transactions: []*Tx{
			tx,
		},
	}

	chain := &blockchain{
		LastHash:   "2",
		Difficulty: 3,
		Height:     2,
	}

	t.Run("Should append peer block to blockchain", func(t *testing.T) {
		chain.addPeerBlock(&peerBlock)

		utils.FailIfDifferent(t, peerBlock.Hash, chain.LastHash)
		utils.FailIfDifferent(t, peerBlock.Difficulty, chain.Difficulty)
		utils.FailIfDifferent(t, peerBlock.Height, chain.Height)

	})

	t.Run("Should remove peerBlock transactions from mempool", func(t *testing.T) {
		onceMP = *new(sync.Once)
		pool := getMP()
		pool.addTx(tx)
		chain.addPeerBlock(&peerBlock)

		utils.FailIfDifferent(t, 0, len(pool.txs))

	})

}

func TestReplace(t *testing.T) {
	blocks := []*Block{
		{Hash: "3", PreviousHash: "2", Difficulty: 4, Height: 3},
		{Hash: "2", PreviousHash: "1"},
		{Hash: "1", PreviousHash: ""},
	}

	chain := &blockchain{
		LastHash:   "42",
		Difficulty: 2,
		Height:     1,
	}

	t.Run("Should replace blockchain", func(t *testing.T) {
		chain.replace(blocks)
		lastBlock := blocks[0]

		utils.FailIfDifferent(t, lastBlock.Hash, chain.LastHash)
		utils.FailIfDifferent(t, lastBlock.Difficulty, chain.Difficulty)
		utils.FailIfDifferent(t, lastBlock.Height, chain.Height)

	})

}

func Test(t *testing.T) {

	t.Run("", func(t *testing.T) {

	})

}
