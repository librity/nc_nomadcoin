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

type TxInput struct {
	Owner  string `json:"owner"`
	Amount uint   `json:"amount"`
}

type TxOutput struct {
	Owner  string `json:"owner"`
	Amount uint   `json:"amount"`
}

func (b *blockchain) TxOutputsFrom(address string) []*TxOutput {
	var outputsFrom []*TxOutput
	outputs := b.txOutputs()

	for _, output := range outputs {
		if output.Owner == address {
			outputsFrom = append(outputsFrom, output)
		}
	}

	return outputsFrom
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
	oldOutputs := Get().TxOutputsFrom(from)

	for _, oldOutput := range oldOutputs {
		if total > amount {
			break
		}

		newInput := &TxInput{oldOutput.Owner, oldOutput.Amount}
		inputs = append(inputs, newInput)
		total += newInput.Amount
	}

	return inputs, total
}

func generateOutputs(from string, to string, amount uint, total uint) []*TxOutput {
	outputs := []*TxOutput{
		{to, amount},
	}
	change := total - amount
	if change > 0 {
		changeOutput := &TxOutput{from, change}
		outputs = append(outputs, changeOutput)
	}

	return outputs
}

func exceedesBalance(from string, amount uint) bool {
	fromBalance := Get().BalanceOf(from)

	return fromBalance < amount
}

func makeCoinbaseTx(address string) *Tx {
	inputs := []*TxInput{
		{coinbaseCode, minerReward},
	}
	outputs := []*TxOutput{
		{address, minerReward},
	}
	tx := newTx(inputs, outputs)

	return tx
}

func (t *Tx) setId() {
	t.Id = utils.HexHash(t)
}

func (b *blockchain) txOutputs() []*TxOutput {
	var outputs []*TxOutput
	blocks := b.Blocks()

	for _, block := range blocks {
		for _, tx := range block.Transactions {
			outputs = append(outputs, tx.Outputs...)
		}
	}

	return outputs
}
