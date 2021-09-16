package utils

import "testing"

func TestHexHash(t *testing.T) {
	testS := struct{ Test string }{Test: "I am a test struct."}
	hash := HexHash(testS)
	t.Logf(hash)
}
