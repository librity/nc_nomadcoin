package blockchain

func GetAddresses() []string {
	return getAddresses(getBC())
}

func getAddresses(chain *blockchain) []string {
	txs := getTxs(chain)
	addresses := make(map[string]bool)

	for _, tx := range txs {
		for _, output := range tx.Outputs {
			addresses[output.Address] = true
		}
	}

	var addressList []string
	for address := range addresses {
		addressList = append(addressList, address)
	}
	return addressList
}
