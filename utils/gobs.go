package utils

import (
	"bytes"
	"encoding/gob"
)

// FromGob encodes an interface pointer to golang byte slice (gob).
func ToGob(i interface{}) []byte {
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(i)
	PanicError(err)

	return buffer.Bytes()
}

// FromGob decodes a golang byte slice (gob) to an interface pointer.
func FromGob(target interface{}, encoded []byte) {
	buffer := bytes.NewReader(encoded)
	decoder := gob.NewDecoder(buffer)
	err := decoder.Decode(target)
	PanicError(err)
}
