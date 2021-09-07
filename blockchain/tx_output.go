package blockchain

type TxOutput struct {
	Owner  string `json:"owner"`
	Amount uint   `json:"amount"`
}

func newTxOutput(address string, amount uint) *TxOutput {
	newOutput := &TxOutput{
		Owner:  address,
		Amount: amount,
	}

	return newOutput
}
