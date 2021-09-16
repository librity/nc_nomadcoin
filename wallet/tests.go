package wallet

func makeTestWallet(keyHex string) *wallet {
	key := keyFromHex(keyHex)
	testWallet := newWallet(key)

	return testWallet
}
