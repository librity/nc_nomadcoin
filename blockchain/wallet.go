package blockchain

func BalanceOf(address string) uint {
	outputs := UnspentTxOutputsFrom(address)

	return SumOverBalance(outputs)
}

func SumOverBalance(outputs []*UnspentTxOutput) uint {
	balance := uint(0)

	for _, output := range outputs {
		balance += output.Amount
	}

	return balance
}
