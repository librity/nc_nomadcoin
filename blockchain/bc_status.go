package blockchain

type BCStatus struct {
	LastHash         string `json:"lastHash"`
	Height           int    `json:"height"`
	CurrentDificulty int    `json:"currentDifficulty"`
}

func Status() *BCStatus {
	chain := getBC()
	status := newBCStatus(chain)

	return status
}

func newBCStatus(b *blockchain) *BCStatus {
	bcStatus := &BCStatus{
		LastHash:         b.LastHash,
		Height:           b.Height,
		CurrentDificulty: b.Dificulty,
	}

	return bcStatus
}
