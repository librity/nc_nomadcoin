package blockchain

import (
	"sync"

	"github.com/librity/nc_nomadcoin/wallet"
)

type mempool struct {
	transactions []*Tx
	m            sync.Mutex
}

var (
	mp     *mempool
	onceMP sync.Once
)

func AddTx(to string, amount uint) (*Tx, error) {
	from := wallet.GetAddress()
	tx, err := makeTx(from, to, amount)
	if err != nil {
		return nil, err
	}

	addTxToMP(tx)
	return tx, nil
}

func AddPeerTx(peerTx *Tx) {
	// TODO: Validate tx

	addTxToMP(peerTx)
}

func getMP() *mempool {
	if mp == nil {
		onceMP.Do(initializeMP)
	}

	return mp
}

func initializeMP() {
	mp = &mempool{}
}

func addTxToMP(tx *Tx) {
	pool := getMP()
	pool.m.Lock()
	defer pool.m.Unlock()

	pool.transactions = append(pool.transactions, tx)
}

func popAllFromMP() []*Tx {
	pool := getMP()
	pool.m.Lock()
	defer pool.m.Unlock()

	miner := wallet.GetAddress()
	coinbase := makeCoinbaseTx(miner)
	txs := pool.transactions
	txs = append(txs, coinbase)
	pool.clear()

	return txs
}

func (m *mempool) clear() {
	m.transactions = nil
}

func outputIsOnMP(unspentOutput *UnspTxOutput) bool {
	for _, transaction := range getMP().transactions {
		for _, input := range transaction.Inputs {
			sameTxId := input.TxId == unspentOutput.TxId
			sameIndex := input.OutputIndex == unspentOutput.Index
			outputIsOnMempool := sameTxId && sameIndex

			if outputIsOnMempool {
				return true
			}
		}
	}

	return false
}
