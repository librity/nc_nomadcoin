package blockchain

type mempool struct {
	Transactions []*Tx `json:"transactions"`
}

var (
	Mempool = &mempool{}
)

func (m *mempool) AddTx(to string, amount uint) error {
	tx, err := makeTx("lior", to, amount)
	if err != nil {
		return err
	}

	m.Transactions = append(m.Transactions, tx)
	return nil
}

func (m *mempool) popAll() []*Tx {
	coinbase := makeCoinbaseTx("lior")
	txs := m.Transactions
	txs = append(txs, coinbase)
	m.Transactions = nil

	return txs
}
