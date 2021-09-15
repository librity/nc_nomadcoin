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

	getMP().addTx(tx)
	return tx, nil
}

func AddPeerTx(peerTx *Tx) {
	// TODO: Validate transaction

	getMP().addTx(peerTx)
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

func (m *mempool) addTx(tx *Tx) {
	m.m.Lock()
	defer m.m.Unlock()

	m.transactions = append(m.transactions, tx)
}

func (m *mempool) popAll() []*Tx {
	m.m.Lock()
	defer m.m.Unlock()

	miner := wallet.GetAddress()
	coinbase := makeCoinbaseTx(miner)
	txs := m.transactions
	txs = append(txs, coinbase)
	m.clear()

	return txs
}

func (m *mempool) removeConfirmedTxs(peerBlock *Block) {
	m.m.Lock()
	defer m.m.Unlock()

	confirmedTxs := peerBlock.Transactions
	for _, confirmedTx := range confirmedTxs {
		isInMP, index := m.hasTx(confirmedTx.Id)
		if isInMP {
			m.removeTx(index)
		}
	}
}

func (m *mempool) hasTx(targetId string) (bool, int) {
	for index, tx := range m.transactions {
		if tx.Id == targetId {
			return true, index
		}
	}

	return false, -1
}

func (m *mempool) removeTx(index int) {
	m.transactions = append(m.transactions[:index], m.transactions[index+1:]...)
}

func (m *mempool) clear() {
	m.transactions = nil
}
