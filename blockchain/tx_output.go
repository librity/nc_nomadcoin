package blockchain

type TxOutput struct {
	Address string `json:"address"`
	Amount  uint   `json:"amount"`
}

func newTxOutput(address string, amount uint) *TxOutput {
	newOutput := &TxOutput{
		Address: address,
		Amount:  amount,
	}

	return newOutput
}
