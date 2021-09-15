package blockchain

type MPStatus struct {
	Size         int   `json:"size"`
	Transactions []*Tx `json:"transactions"`
}

func MempoolStatus() *MPStatus {
	pool := getMP()
	status := newMPStatus(pool)

	return status
}

func newMPStatus(m *mempool) *MPStatus {
	MPStatus := &MPStatus{
		Size:         len(m.transactions),
		Transactions: m.transactions,
	}

	return MPStatus
}
