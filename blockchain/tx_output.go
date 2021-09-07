package blockchain

type TxOutput struct {
	Owner  string `json:"owner"`
	Amount uint   `json:"amount"`
}

type UnspentTxOutput struct {
	TxId   string `json:"transactionId"`
	Index  uint   `json:"index"`
	Amount uint   `json:"amount"`
}

func (b *blockchain) UnspentTxOutputsFrom(address string) []*TxOutput {
	return nil
}
