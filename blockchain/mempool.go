package blockchain

import (
	"sync"

	"github.com/librity/nc_nomadcoin/wallet"
)

type mempool struct {
	txs map[string]*Tx
	m   sync.Mutex
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
	mp = &mempool{
		txs: make(map[string]*Tx),
	}
}

func (m *mempool) addTx(tx *Tx) {
	m.m.Lock()
	defer m.m.Unlock()

	m.txs[tx.Id] = tx
}

func (m *mempool) popAll() []*Tx {
	m.m.Lock()
	defer m.m.Unlock()

	miner := wallet.GetAddress()
	coinbase := makeCoinbaseTx(miner)
	txs := m.getTxs()
	txs = append(txs, coinbase)
	m.clearTxs()

	return txs
}

func (m *mempool) getTxs() []*Tx {
	txs := []*Tx{}
	for _, tx := range m.txs {
		txs = append(txs, tx)
	}

	return txs
}

func (m *mempool) removeConfirmedTxs(peerBlock *Block) {
	m.m.Lock()
	defer m.m.Unlock()

	confirmedTxs := peerBlock.Txs
	for _, confirmedTx := range confirmedTxs {
		_, txIsInMP := m.txs[confirmedTx.Id]
		if txIsInMP {
			m.removeTx(confirmedTx.Id)
		}
	}
}

func (m *mempool) removeTx(taretId string) {
	delete(m.txs, taretId)
}

func (m *mempool) clearTxs() {
	m.txs = make(map[string]*Tx)
}
