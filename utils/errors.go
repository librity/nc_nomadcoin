package utils

import "log"

var (
	panicFn = log.Panic
)

// PanicError panics the error unless it is nill.
func PanicError(err error) {
	if err != nil {
		panicFn(err)
	}
}
