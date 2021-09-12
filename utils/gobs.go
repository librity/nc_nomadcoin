package utils

import (
	"bytes"
	"encoding/gob"
)

func ToGob(i interface{}) []byte {
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(i)
	PanicError(err)

	return buffer.Bytes()
}

func FromGob(target interface{}, encoded []byte) {
	buffer := bytes.NewReader(encoded)
	decoder := gob.NewDecoder(buffer)
	err := decoder.Decode(target)
	PanicError(err)
}
