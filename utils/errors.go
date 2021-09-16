package utils

import "log"

// PanicError panics the error unless it is nill.
func PanicError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
