package blockchain

func makeFakeTx(id string) *Tx {
	fakeTx := &Tx{
		Id:        id,
		Timestamp: 0,
		Inputs:    []*TxInput{},
		Outputs:   []*TxOutput{},
	}

	return fakeTx
}
