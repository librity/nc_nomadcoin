package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestToJSON(t *testing.T) {
	type testStruct struct{ Test string }
	test := testStruct{Test: "I am a test struct."}

	t.Run("Should return a slice of bytes", func(t *testing.T) {
		result := ToJSON(test)
		kind := reflect.TypeOf(result).Kind()

		if kind != reflect.Slice {
			t.Errorf("Should return a slice of bytes, got %s", kind)
		}
	})

	t.Run("Should return a decodable gob", func(t *testing.T) {
		encoded := ToJSON(test)
		decoded := testStruct{}
		err := json.Unmarshal(encoded, &decoded)

		if err != nil {
			t.Error("Unable to decode Gob.")
		}

		if decoded != test {
			t.Error("Decoded gob doesn't match original object.")
		}
	})

}

func ExampleToJSON() {
	test := struct{ Test string }{Test: "I am a test struct."}
	encoded := ToJSON(test)
	decoded := struct{ Test string }{}
	json.Unmarshal(encoded, &decoded)
	fmt.Println(decoded)
	// Output: {I am a test struct.}
}

func TestFromJSON(t *testing.T) {
	type testStruct struct{ Test string }
	test := testStruct{Test: "I am a test struct."}
	encoded := []byte(`{"Test": "I am a test struct."}`)

	t.Run("Should decode the json into the original object", func(t *testing.T) {
		decoded := testStruct{}
		FromJSON(encoded, &decoded)

		if !reflect.DeepEqual(test, decoded) {
			t.Errorf("Should return a %s, got %s", reflect.TypeOf(test), reflect.TypeOf(decoded))
		}

		if decoded != test {
			t.Error("Decoded gob doesn't match original object.")
		}
	})

}

func ExampleFromJSON() {
	test := struct{ Test string }{Test: "I am a test struct."}
	encoded, _ := json.Marshal(test)
	decoded := struct{ Test string }{}
	FromJSON(encoded, &decoded)
	fmt.Println(decoded)
	// Output: {I am a test struct.}
}
