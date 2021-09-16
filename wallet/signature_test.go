package wallet

import (
	"encoding/hex"
	"testing"
)

const (
	testKeyHex    = "30770201010420876f9b0868241e52f7ff0ba997d40f09062442890613443990507ff9066b6b12a00a06082a8648ce3d030107a14403420004c544a8f9319ad8e2688a4c66c94cb6ed434b86058ad1489d4d152ff6fc47943718ce246f35ef03c8b3dce4e39bf409a3016830a5b1eef6c441453951586b2903"
	randomHash    = "b2f855e131720ae951ba70a671ea6deb60702bc38932986f53be7ab592065699"
	testSignature = "11a6f5353027b8af5f1b67de2543cb37329553ef4eb595f7b22eb614d499a0218a2dd1380c57d2267c4ad4f22af584b321fbbfc1df8705ac87821caea98715b2"
)

var (
	testWallet = makeTestWallet(testKeyHex)
)

func makeTestWallet(keyHex string) *wallet {
	key := keyFromHex(keyHex)
	testWallet := newWallet(key)

	return testWallet
}

func TestHexSign(t *testing.T) {
	signature := HexSign(randomHash, testWallet)
	_, err := hex.DecodeString(signature)
	if err != nil {
		t.Errorf("Should return a hex-encoded string, got \"%s\"", signature)
	}
}

func TestVerify(t *testing.T) {
	signature := HexSign(randomHash, testWallet)
	_, err := hex.DecodeString(signature)
	if err != nil {
		t.Errorf("Should return a hex-encoded string, got \"%s\"", signature)
	}
}
