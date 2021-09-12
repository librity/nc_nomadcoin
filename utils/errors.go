package utils

import "log"

func PanicError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
