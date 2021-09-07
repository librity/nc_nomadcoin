package blockchain

type TxInput struct {
	TxId  string `json:"transactionId"`
	Index uint   `json:"index"`
	Owner string `json:"owner"`
}
