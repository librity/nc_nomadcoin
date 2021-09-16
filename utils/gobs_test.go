package utils

import (
	"bytes"
	"encoding/gob"
	"reflect"
	"testing"
)

func TestToGob(t *testing.T) {
	test := struct{ Test string }{Test: "I am a test struct."}

	t.Run("Should return a slice of bytes", func(t *testing.T) {
		result := ToGob(test)
		kind := reflect.TypeOf(result).Kind()

		if kind != reflect.Slice {
			t.Errorf("Should return a slice of bytes, got %s", kind)
		}
	})

	t.Run("Should return a decodable gob", func(t *testing.T) {
		encoded := ToGob(test)
		decoded := struct{ Test string }{}
		err := gob.NewDecoder(bytes.NewReader(encoded)).Decode(&decoded)

		if err != nil {
			t.Error("Unable to decode Gob.")
		}

		if decoded != test {
			t.Error("Decoded gob doesn't match original object.")
		}
	})

}
