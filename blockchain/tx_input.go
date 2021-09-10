package blockchain

const (
	coinbaseCode = "COINBASE"
)

type TxInput struct {
	TxId      string `json:"transactionId"`
	Index     uint   `json:"index"`
	Signature string `json:"signature"`
}

func newTxInput(unspentOutput *UnspTxOutput, signature string) *TxInput {
	newInput := &TxInput{
		TxId:      unspentOutput.TxId,
		Index:     unspentOutput.Index,
		Signature: signature,
	}

	return newInput
}

func newCoinbaseTxInput() *TxInput {
	coinbaseInput := &TxInput{
		TxId:      "",
		Index:     0,
		Signature: coinbaseCode,
	}

	return coinbaseInput
}
