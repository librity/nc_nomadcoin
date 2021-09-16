package utils

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestHexHash(t *testing.T) {
	expectedHash := "4866e29dd676294d1920a286d872b935f7a5d7733a1b9aeecb005a882229d6f3"
	test := struct{ Test string }{Test: "I am a test struct."}

	t.Run("Hash is deterministic", func(t *testing.T) {
		firstCall := HexHash(test)
		t.Logf(firstCall)

		if firstCall != expectedHash {
			t.Errorf("Expectd \"%s\", got \"%s\"", expectedHash, firstCall)
		}

		secondCall := HexHash(test)
		if secondCall != firstCall {
			t.Errorf("Expectd \"%s\", got \"%s\"", firstCall, secondCall)
		}
	})

	t.Run("Hash is properly hex encoded", func(t *testing.T) {
		hash := HexHash(test)
		_, err := hex.DecodeString(hash)
		if err != nil {
			t.Error("Hash should be hex encoded")
		}
	})
}

func ExampleHexHash() {
	test := struct{ Test string }{Test: "I am a test struct."}
	hash := HexHash(test)
	fmt.Println(hash)
	// Output: 4866e29dd676294d1920a286d872b935f7a5d7733a1b9aeecb005a882229d6f3
}
