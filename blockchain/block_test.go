package blockchain

import (
	"reflect"
	"testing"

	"github.com/librity/nc_nomadcoin/utils"
)

func makeFakeBlock() *Block {
	fakeBlock := &Block{
		Height:       1,
		PreviousHash: "",
		Hash:         utils.RandomHash(),
		Difficulty:   1,
		NOnce:        0,
	}

	return fakeBlock
}

func TestCreateBlock(t *testing.T) {
	storage = fakeStorageLayer{}

	t.Run("Should return a *Block", func(t *testing.T) {
		block := createBlock("GENESIS", 1, 1)
		bType := reflect.TypeOf(block)
		expectedType := reflect.TypeOf(&Block{})

		if bType != expectedType {
			t.Errorf("Expected %v, got %v", expectedType, bType)
		}
	})

	t.Run("Should load transactions from the mempool", func(t *testing.T) {
		txId := "test"
		fakeTx := makeFakeTx(txId)
		getMP().txs[txId] = fakeTx
		block := createBlock("GENESIS", 1, 1)

		blockTx := block.Transactions[0]
		if blockTx != fakeTx {
			t.Errorf("Expected %v, got %v", fakeTx, blockTx)
		}

		if blockTx.Id != txId {
			t.Errorf("Expected %v, got %v", txId, blockTx.Id)
		}

		mpSize := len(getMP().txs)
		if mpSize != 0 {
			t.Error("Mined block should clear Mempool transactions.")
		}
	})

}
