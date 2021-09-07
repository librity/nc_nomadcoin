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
	var spentTxs map[string]bool

	for _, block := range b.Blocks() {
		for _, tx := range block.Transactions {
			for _, input := range tx.Inputs {
				if input.Owner == address {
					spentTxs[input.TxId] = true
				}
			}

			for index, output := range tx.Outputs {
				if output.Owner != address {
					continue
				}

				_, spentInput := spentTxs[tx.Id]
				if spentInput == false {
					unspentOutput := &UnspentTxOutput{
						TxId:   tx.Id,
						Index:  uint(index),
						Amount: output.Amount,
					}

					unspentOutputs = append(unspentOutputs, unspentOutput)
				}
			}
		}
	}

	return unspentOutputs
}

func newTxOutput(address string, amount uint) *TxOutput {
	newOutput := &TxOutput{
		Owner:  address,
		Amount: amount,
	}

	return newOutput
}
