package utils

import "encoding/json"

// ToJSON encodes an interface to a json byte slice.
func ToJSON(i interface{}) []byte {
	bytes, err := json.Marshal(i)
	PanicError(err)

	return bytes
}

// FromJSON decodes a json byte slice to an interface pointer.
func FromJSON(bytes []byte, i interface{}) {
	err := json.Unmarshal(bytes, i)
	PanicError(err)
}
