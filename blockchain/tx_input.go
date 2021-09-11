package blockchain

const (
	coinbaseCode = "COINBASE"
)

type TxInput struct {
	TxId        string `json:"transactionId"`
	OutputIndex uint   `json:"outputIndex"`
	Signature   string `json:"signature"`
}

func newTxInput(unspentOutput *UnspTxOutput, signature string) *TxInput {
	newInput := &TxInput{
		TxId:        unspentOutput.TxId,
		OutputIndex: unspentOutput.Index,
		Signature:   signature,
	}

	return newInput
}

func newCoinbaseTxInput() *TxInput {
	coinbaseInput := &TxInput{
		TxId:        "",
		OutputIndex: 0,
		Signature:   coinbaseCode,
	}

	return coinbaseInput
}
