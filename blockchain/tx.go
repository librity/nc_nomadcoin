package blockchain

import (
	"errors"

	"github.com/librity/nc_nomadcoin/utils"
)

const (
	minerReward = 50
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
		newCoinbaseTxInput(),
	}
	outputs := []*TxOutput{
		{address, minerReward},
	}
	tx := newTx(inputs, outputs)

	return tx
}

func makeTx(from string, to string, amount uint) (*Tx, error) {
	if exceedesBalance(from, amount) {
		return nil, ErrNotEnoughMoney
	}

	inputs, total := generateInputs(from, amount)
	outputs := generateOutputs(from, to, amount, total)
	tx := newTx(inputs, outputs)

	return tx, nil
}

func generateInputs(from string, amount uint) ([]*TxInput, uint) {
	var inputs []*TxInput
	total := uint(0)
	unspentOutputs := Get().UnspentTxOutputsFrom(from)

	for _, unspentOutput := range unspentOutputs {
		if total > amount {
			break
		}

		newInput := newTxInput(unspentOutput, from)
		inputs = append(inputs, newInput)
		total += unspentOutput.Amount
	}

	return inputs, total
}

func generateOutputs(from string, to string, amount uint, total uint) []*TxOutput {
	toOutput := newTxOutput(to, amount)
	outputs := []*TxOutput{toOutput}
	change := total - amount
	if change > 0 {
		changeOutput := &TxOutput{
			Owner:  from,
			Amount: change,
		}
		outputs = append(outputs, changeOutput)
	}

	return outputs
}

func exceedesBalance(from string, amount uint) bool {
	fromBalance := Get().BalanceOf(from)
	return fromBalance < amount
}

func (t *Tx) setId() {
	t.Id = utils.HexHash(t)
}
