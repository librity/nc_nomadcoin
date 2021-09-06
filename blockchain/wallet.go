package blockchain

func (b *blockchain) BalanceOf(address string) uint {
	outputs := b.TxOutputsFrom(address)

	return SumOverBalance(outputs)
}

func SumOverBalance(outputs []*TxOutput) uint {
	balance := uint(0)

	for _, output := range outputs {
		balance += output.Amount
	}

	return balance
}
