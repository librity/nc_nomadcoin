package blockchain

type TxOutput struct {
	Owner  string `json:"owner"`
	Amount uint   `json:"amount"`
}

type UnspentTxOutput struct {
	TxId   string `json:"transactionId"`
	Index  uint   `json:"index"`
	Amount uint   `json:"amount"`
}

func (b *blockchain) UnspentTxOutputsFrom(address string) []*UnspentTxOutput {
	var unspentOutputs []*UnspentTxOutput
	spentTxs := make(map[string]bool)

	for _, block := range b.Blocks() {
		for _, tx := range block.Transactions {
			setSpentTxs(spentTxs, tx, address)
			setUnspentOutputs(&unspentOutputs, spentTxs, tx, address)
		}
	}

	return unspentOutputs
}

func setSpentTxs(spentTxs map[string]bool, tx *Tx, address string) {
	for _, input := range tx.Inputs {
		if input.Owner == address {
			spentTxs[input.TxId] = true
		}
	}
}

func setUnspentOutputs(
	unspentOutputs *[]*UnspentTxOutput,
	spentTxs map[string]bool,
	tx *Tx,
	address string) {
	for index, output := range tx.Outputs {
		if output.Owner != address {
			continue
		}

		_, spentInput := spentTxs[tx.Id]
		if !spentInput {
			unspentOutput := newUnspentTxOutput(
				tx.Id, uint(index), output.Amount)

			*unspentOutputs = append(*unspentOutputs, unspentOutput)
		}
	}
}

func newTxOutput(address string, amount uint) *TxOutput {
	newOutput := &TxOutput{
		Owner:  address,
		Amount: amount,
	}

	return newOutput
}

func newUnspentTxOutput(txId string, index uint, amount uint) *UnspentTxOutput {
	newUnspentOutput := &UnspentTxOutput{
		TxId:   txId,
		Index:  index,
		Amount: amount,
	}

	return newUnspentOutput
}
