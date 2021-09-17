package wallet

const (
	testKeyHex = "30770201010420876f9b0868241e52f7ff0ba997d40f09062442890613443990507ff9066b6b12a00a06082a8648ce3d030107a14403420004c544a8f9319ad8e2688a4c66c94cb6ed434b86058ad1489d4d152ff6fc47943718ce246f35ef03c8b3dce4e39bf409a3016830a5b1eef6c441453951586b2903"
)

var (
	testWallet = makeTestWallet(testKeyHex)
)

func makeTestWallet(keyHex string) *wallet {
	key := keyFromHex(keyHex)
	testWallet := newWallet(key)

	return testWallet
}
