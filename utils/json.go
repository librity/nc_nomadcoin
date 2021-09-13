package utils

import "encoding/json"

func ToJSON(payload interface{}) []byte {
	bytes, err := json.Marshal(payload)
	PanicError(err)

	return bytes
}
