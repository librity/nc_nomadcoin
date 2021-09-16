package wallet

import (
	"encoding/hex"
	"testing"
)

const (
	testKeyHex    = "30770201010420876f9b0868241e52f7ff0ba997d40f09062442890613443990507ff9066b6b12a00a06082a8648ce3d030107a14403420004c544a8f9319ad8e2688a4c66c94cb6ed434b86058ad1489d4d152ff6fc47943718ce246f35ef03c8b3dce4e39bf409a3016830a5b1eef6c441453951586b2903"
	testHash      = "b2f855e131720ae951ba70a671ea6deb60702bc38932986f53be7ab592065699"
	testSignature = "11a6f5353027b8af5f1b67de2543cb37329553ef4eb595f7b22eb614d499a0218a2dd1380c57d2267c4ad4f22af584b321fbbfc1df8705ac87821caea98715b2"
	testAddress   = "c544a8f9319ad8e2688a4c66c94cb6ed434b86058ad1489d4d152ff6fc47943718ce246f35ef03c8b3dce4e39bf409a3016830a5b1eef6c441453951586b2903"

	badHash      = "0324c89217d5215b76292bf78480d466bdd3963d051c9e70d0844f13cbbb9408"
	badSignature = "439310e083e6d45ad4f4cad3d59f43879c4392ab0a237069f1dc85129a384f01f897ddf2b3e7b4a7a6137586088574c060d9c14fbf0bdd4f30d0a965e36d33fc"
	badAddress   = "aa369c06c0b0ee60f6effdd87777d90181e7c414173e07316680190ead3f1e413255c2830989b73c195b0e0bb4dd68150a675435965ec694bb2f0c7155f0878d"
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
	signature := HexSign(testHash, testWallet)
	_, err := hex.DecodeString(signature)
	if err != nil {
		t.Errorf("Should return a hex-encoded string, got \"%s\"", signature)
	}
}

func TestVerify(t *testing.T) {
	type testCase struct {
		hash     string
		sign     string
		address  string
		expected bool
	}

	testCases := []testCase{
		{
			hash:     testHash,
			sign:     testSignature,
			address:  testAddress,
			expected: true},
		{
			hash:     badHash,
			sign:     testSignature,
			address:  testAddress,
			expected: false},
		{
			hash:     testHash,
			sign:     badSignature,
			address:  testAddress,
			expected: false},
		{
			hash:     testHash,
			sign:     testSignature,
			address:  badAddress,
			expected: false},
	}

	for _, tc := range testCases {
		result := Verify(tc.hash, tc.sign, tc.address)
		if result != tc.expected {
			t.Error("Should return true with valid hash, signature and address.")
		}
	}
}
