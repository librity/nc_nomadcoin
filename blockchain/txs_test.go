package blockchain

import (
	"testing"

	"github.com/librity/nc_nomadcoin/utils"
)

func TestFindTx(t *testing.T) {

	t.Run("Should return ErrTxNotFound when Tx doesn't exist", func(t *testing.T) {
		fakeBlock := makeFakeBlock()
		storage = fakeStorageLayer{
			fakeLoadBlock: func() []byte { return utils.ToGob(fakeBlock) },
		}
		chain := &blockchain{LastHash: fakeBlock.Hash}

		tx, err := findTx(chain, "test")
		if err != ErrTxNotFound {
			t.Errorf("Expected %v, got %v", ErrTxNotFound, err)
		}

		if tx != nil {
			t.Errorf("Expected %v, got %v", nil, tx)
		}
	})

	t.Run("Should return the correct *Tx", func(t *testing.T) {
		blockControl := 0
		storage = fakeStorageLayer{
			fakeLoadBlock: func() []byte {
				fakeBlocks := []*Block{
					{Hash: "3", PreviousHash: "2"},
					{Hash: "2", PreviousHash: "1",
						Transactions: []*Tx{makeFakeTx("test")}},
					{Hash: "1", PreviousHash: ""},
				}

				chosenBlock := fakeBlocks[blockControl]
				blockControl++
				return utils.ToGob(chosenBlock)
			},
		}
		chain := &blockchain{LastHash: "3"}

		tx, err := findTx(chain, "test")
		if err != nil {
			t.Errorf("Expected %v, got %v", nil, err)
		}

		if tx.Id != "test" {
			t.Errorf("Expected %v, got %v", "test", tx.Id)
		}
	})

}

func Test(t *testing.T) {

	t.Run("", func(t *testing.T) {

	})

}
