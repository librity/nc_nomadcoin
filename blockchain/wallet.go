package blockchain

func BalanceOf(address string) uint {
	outputs := UnspTxOutputsFrom(address)

	return SumOverBalance(outputs)
}

func SumOverBalance(outputs []*UnspTxOutput) uint {
	balance := uint(0)

	for _, output := range outputs {
		balance += output.Amount
	}

	return balance
}
