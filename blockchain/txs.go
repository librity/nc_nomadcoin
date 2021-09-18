package blockchain

import "errors"

var (
	ErrTxNotFound = errors.New("transaction not found")
)

func GetTxs() []*Tx {
	return getTxs(getBC())
}

func getTxs(chain *blockchain) []*Tx {
	txs := []*Tx{}

	for _, block := range getBlocks(chain) {
		txs = append(txs, block.Transactions...)
	}

	return txs
}

func FindTx(targetId string) (*Tx, error) {
	return findTx(getBC(), targetId)
}

func findTx(chain *blockchain, targetId string) (*Tx, error) {
	for _, tx := range getTxs(chain) {
		if tx.Id == targetId {
			return tx, nil
		}
	}

	return nil, ErrTxNotFound
}
