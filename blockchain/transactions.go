package blockchain

import "github.com/librity/nc_nomadcoin/utils"

const (
	coinbaseCode = "COINBASE"
	minerReward  = 50
)

// Transaction
type Tx struct {
	Id        string      `json:"id"`
	Timestamp int         `json:"timestamp"`
	TxInputs  []*TxInput  `json:"txInputs"`
	TxOutputs []*TxOutput `json:"txOutputs"`
}

type TxInput struct {
	Owner  string `json:"owner"`
	Amount int    `json:"amount"`
}

type TxOutput struct {
	Owner  string `json:"owner"`
	Amount int    `json:"amount"`
}

func makeCoinbaseTx(address string) *Tx {
	txInputs := []*TxInput{
		{coinbaseCode, minerReward},
	}
	txOutputs := []*TxOutput{
		{address, minerReward},
	}

	tx := Tx{
		Id:        "",
		Timestamp: utils.Now(),
		TxInputs:  txInputs,
		TxOutputs: txOutputs,
	}
	tx.setId()

	return &tx
}

func (t *Tx) setId() {
	t.Id = utils.HexHash(t)
}
