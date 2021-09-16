package blockchain

type UnspTxOutput struct {
	TxId   string `json:"transactionId"`
	Index  uint   `json:"index"`
	Amount uint   `json:"amount"`
}

func UnspTxOutputsFrom(address string) []*UnspTxOutput {
	allTxs := GetTxs()
	referencedTxs := filterReferencedTxs(&allTxs, address)
	unspentOutputs := makeUnspOutputs(&allTxs, referencedTxs, address)

	return unspentOutputs
}

func filterReferencedTxs(txs *[]*Tx, address string) map[string]bool {
	referencedTxs := make(map[string]bool)

	for _, tx := range *txs {
		for _, input := range tx.Inputs {
			creatorOutput, err := findCreatorOutput(input)
			if err != nil {
				continue
			}

			if creatorOutput.Address != address {
				continue
			}

			referencedTxs[input.TxId] = true
		}
	}

	return referencedTxs
}

func makeUnspOutputs(txs *[]*Tx, referencedTxs map[string]bool, address string) []*UnspTxOutput {
	var unspentOutputs []*UnspTxOutput

	for _, tx := range *txs {
		for index, output := range tx.Outputs {
			if output.Address != address {
				continue
			}

			_, isSpentInput := referencedTxs[tx.Id]
			if isSpentInput {
				continue
			}

			unspentOutput := newUnspTxOutput(tx.Id, uint(index), output.Amount)
			if outputIsOnMP(unspentOutput) {
				continue
			}

			unspentOutputs = append(unspentOutputs, unspentOutput)
		}
	}

	return unspentOutputs
}

func outputIsOnMP(unspentOutput *UnspTxOutput) bool {
	for _, transaction := range getMP().txs {
		for _, input := range transaction.Inputs {
			sameTxId := input.TxId == unspentOutput.TxId
			sameIndex := input.OutputIndex == unspentOutput.Index
			outputIsOnMempool := sameTxId && sameIndex

			if outputIsOnMempool {
				return true
			}
		}
	}

	return false
}

func newUnspTxOutput(txId string, index, amount uint) *UnspTxOutput {
	newUnspentOutput := &UnspTxOutput{
		TxId:   txId,
		Index:  index,
		Amount: amount,
	}

	return newUnspentOutput
}
