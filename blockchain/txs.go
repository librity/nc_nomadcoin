package blockchain

import "errors"

var (
	ErrTxNotFound = errors.New("transaction not found")
)

func GetTxs() []*Tx {
	txs := []*Tx{}

	for _, block := range GetBlocks() {
		txs = append(txs, block.Transactions...)
	}

	return txs
}

func FindTx(targetId string) (*Tx, error) {
	for _, tx := range GetTxs() {
		if tx.Id == targetId {
			return tx, nil
		}
	}

	return nil, ErrTxNotFound
}
