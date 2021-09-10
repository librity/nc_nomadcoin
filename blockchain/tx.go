package blockchain

import (
	"errors"

	"github.com/librity/nc_nomadcoin/utils"
	"github.com/librity/nc_nomadcoin/wallet"
)

const (
	minerReward = 50
)

var (
	ErrNotEnoughMoney = errors.New("not enough money")
	ErrInvalidTx      = errors.New("transaction signature verification failed")
)

// Transaction
type Tx struct {
	Id        string      `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Inputs    []*TxInput  `json:"inputs"`
	Outputs   []*TxOutput `json:"outputs"`
}

func (t *Tx) setId() {
	t.Id = utils.HexHash(t)
}

func (t *Tx) sign() {
	hash := t.Id

	for _, input := range t.Inputs {
		input.Signature = wallet.HexSign(hash)
	}
}

func (t *Tx) isValid() bool {
	for _, input := range t.Inputs {
		creatorOutput, err := findCreatorOutput(input)
		if err != nil {
			return false
		}

		hash := t.Id
		signHex := input.Signature
		address := creatorOutput.Address
		isValid := wallet.Verify(hash, signHex, address)
		if !isValid {
			return false
		}
	}

	return true
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
		newTxOutput(address, minerReward),
	}
	tx := newTx(inputs, outputs)

	return tx
}

func makeTx(from, to string, amount uint) (*Tx, error) {
	if exceedesBalance(from, amount) {
		return nil, ErrNotEnoughMoney
	}

	inputs, total := makeInputs(from, amount)
	outputs := makeOutputs(from, to, amount, total)
	tx := newTx(inputs, outputs)
	tx.sign()
	if !tx.isValid() {
		return nil, ErrInvalidTx
	}

	return tx, nil
}

func makeInputs(from string, amount uint) ([]*TxInput, uint) {
	var inputs []*TxInput
	total := uint(0)
	unspentOutputs := UnspTxOutputsFrom(from)

	for _, unspentOutput := range unspentOutputs {
		if total >= amount {
			break
		}

		newInput := newTxInput(unspentOutput, from)
		inputs = append(inputs, newInput)
		total += unspentOutput.Amount
	}

	return inputs, total
}

func makeOutputs(from, to string, amount, total uint) []*TxOutput {
	toOutput := newTxOutput(to, amount)
	outputs := []*TxOutput{toOutput}
	change := total - amount

	if change > 0 {
		changeOutput := newTxOutput(from, change)
		outputs = append(outputs, changeOutput)
	}

	return outputs
}

func exceedesBalance(from string, amount uint) bool {
	fromBalance := BalanceOf(from)
	return fromBalance < amount
}
