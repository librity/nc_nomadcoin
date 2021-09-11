package blockchain

import "errors"

var (
	ErrTxOutputNotFound = errors.New("transaction output not found")
	ErrCoinbaseTxInput  = errors.New("coinbase transaction inputs don't reference any output")
)

type TxOutput struct {
	Address string `json:"address"`
	Amount  uint   `json:"amount"`
}

func findCreatorOutput(input *TxInput) (*TxOutput, error) {
	if input.Signature == coinbaseCode {
		return nil, ErrCoinbaseTxInput
	}

	creatorTx, err := FindTx(input.TxId)
	if err != nil {
		return nil, err
	}

	creatorOutput := creatorTx.Outputs[input.OutputIndex]
	if creatorOutput == nil {
		return nil, ErrTxOutputNotFound
	}

	return creatorOutput, nil
}

func newTxOutput(address string, amount uint) *TxOutput {
	newOutput := &TxOutput{
		Address: address,
		Amount:  amount,
	}

	return newOutput
}
