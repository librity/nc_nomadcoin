package blockchain

import (
	"errors"

	"github.com/librity/nc_nomadcoin/utils"
)

const (
	coinbaseCode = "COINBASE"
	minerReward  = 50
)

var (
	ErrNotEnoughMoney = errors.New("not enough money")
)

// Transaction
type Tx struct {
	Id        string      `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Inputs    []*TxInput  `json:"inputs"`
	Outputs   []*TxOutput `json:"outputs"`
}

func newTx(inputs []*TxInput, outputs []*TxOutput) *Tx {
	tx := Tx{
		Id:        "",
		Timestamp: now(),
		Inputs:    inputs,
		Outputs:   outputs,
	}
	tx.setId()

	return &tx
}

func makeCoinbaseTx(address string) *Tx {
	inputs := []*TxInput{
		{"", 0, coinbaseCode},
	}
	outputs := []*TxOutput{
		{address, minerReward},
	}
	tx := newTx(inputs, outputs)

	return tx
}

func makeTx(from string, to string, amount uint) (*Tx, error) {
	return nil, nil
}

func (t *Tx) setId() {
	t.Id = utils.HexHash(t)
}
