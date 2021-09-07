package blockchain

const (
	coinbaseCode = "COINBASE"
)

type TxInput struct {
	TxId  string `json:"transactionId"`
	Index uint   `json:"index"`
	Owner string `json:"owner"`
}

func newTxInput(unspentOutput *UnspentTxOutput, address string) *TxInput {
	newInput := &TxInput{
		TxId:  unspentOutput.TxId,
		Index: unspentOutput.Index,
		Owner: address,
	}

	return newInput
}

func newCoinbaseTxInput() *TxInput {
	coinbaseInput := &TxInput{
		TxId:  "",
		Index: 0,
		Owner: coinbaseCode,
	}

	return coinbaseInput
}
