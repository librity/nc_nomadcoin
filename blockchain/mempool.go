package blockchain

import "github.com/librity/nc_nomadcoin/wallet"

type mempool struct {
	Transactions []*Tx `json:"transactions"`
}

var (
	Mempool = &mempool{}
)

func (m *mempool) AddTx(to string, amount uint) error {
	from := wallet.GetAddress()
	tx, err := makeTx(from, to, amount)
	if err != nil {
		return err
	}

	m.Transactions = append(m.Transactions, tx)
	return nil
}

func (m *mempool) popAll() []*Tx {
	miner := wallet.GetAddress()
	coinbase := makeCoinbaseTx(miner)
	txs := m.Transactions
	txs = append(txs, coinbase)
	m.Transactions = nil

	return txs
}

func isOnMempool(unspentOutput *UnspTxOutput) bool {
	for _, transaction := range Mempool.Transactions {
		for _, input := range transaction.Inputs {
			sameTxId := input.TxId == unspentOutput.TxId
			sameIndex := input.Index == unspentOutput.Index
			outputIsOnMempool := sameTxId && sameIndex

			if outputIsOnMempool {
				return true
			}
		}
	}

	return false
}
