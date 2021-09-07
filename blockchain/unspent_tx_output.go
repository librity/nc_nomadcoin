package blockchain

type UnspentTxOutput struct {
	TxId   string `json:"transactionId"`
	Index  uint   `json:"index"`
	Amount uint   `json:"amount"`
}

func UnspentTxOutputsFrom(address string) []*UnspentTxOutput {
	var unspentOutputs []*UnspentTxOutput
	spentTxs := make(map[string]bool)

	for _, block := range Blocks() {
		for _, tx := range block.Transactions {
			setSpentTxs(spentTxs, tx, address)
			setUnspentOutputs(&unspentOutputs, spentTxs, tx, address)
		}
	}

	return unspentOutputs
}

func setSpentTxs(spentTxs map[string]bool, tx *Tx, address string) {
	for _, input := range tx.Inputs {
		if input.Owner != address {
			continue
		}

		spentTxs[input.TxId] = true
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

		_, isSpentInput := spentTxs[tx.Id]
		if isSpentInput {
			continue
		}

		unspentOutput := newUnspentTxOutput(tx.Id, uint(index), output.Amount)
		if isOnMempool(unspentOutput) {
			continue
		}

		*unspentOutputs = append(*unspentOutputs, unspentOutput)
	}
}

func newUnspentTxOutput(txId string, index uint, amount uint) *UnspentTxOutput {
	newUnspentOutput := &UnspentTxOutput{
		TxId:   txId,
		Index:  index,
		Amount: amount,
	}

	return newUnspentOutput
}
