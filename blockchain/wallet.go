package blockchain

func (b *blockchain) BalanceOf(address string) int {
	outputs := b.TxOutputsFrom(address)

	return SumOverBalance(outputs)
}

func SumOverBalance(outputs []*TxOutput) int {
	balance := 0

	for _, output := range outputs {
		balance += output.Amount
	}

	return balance
}
