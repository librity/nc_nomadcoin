package blockchain

type UnspTxOutput struct {
	TxId   string `json:"transactionId"`
	Index  uint   `json:"index"`
	Amount uint   `json:"amount"`
}

func UnspTxOutputsFrom(address string) []*UnspTxOutput {
	var unspentOutputs []*UnspTxOutput
	creatorTxs := make(map[string]bool)

	for _, block := range GetBlocks() {
		for _, tx := range block.Transactions {
			setCreatorTxs(creatorTxs, tx, address)
			setUnspOutputs(&unspentOutputs, creatorTxs, tx, address)
		}
	}

	return unspentOutputs
}

func setCreatorTxs(creatorTxs map[string]bool, tx *Tx, address string) {
	for _, input := range tx.Inputs {
		if input.Signature != address {
			continue
		}

		creatorTxs[input.TxId] = true
	}
}

func setUnspOutputs(
	unspentOutputs *[]*UnspTxOutput,
	creatorTxs map[string]bool,
	tx *Tx,
	address string) {

	for index, output := range tx.Outputs {
		if output.Address != address {
			continue
		}

		_, isSpentInput := creatorTxs[tx.Id]
		if isSpentInput {
			continue
		}

		unspentOutput := newUnspTxOutput(tx.Id, uint(index), output.Amount)
		if isOnMempool(unspentOutput) {
			continue
		}

		*unspentOutputs = append(*unspentOutputs, unspentOutput)
	}
}

func newUnspTxOutput(txId string, index, amount uint) *UnspTxOutput {
	newUnspentOutput := &UnspTxOutput{
		TxId:   txId,
		Index:  index,
		Amount: amount,
	}

	return newUnspentOutput
}
