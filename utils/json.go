package utils

import "encoding/json"

func ToJSON(i interface{}) []byte {
	bytes, err := json.Marshal(i)
	PanicError(err)

	return bytes
}

func FromJSON(bytes []byte, i interface{}) {
	err := json.Unmarshal(bytes, i)
	PanicError(err)
}
