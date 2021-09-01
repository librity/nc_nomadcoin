package utils

import (
	"log"
	"reflect"
)

func HandleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// Source: https://stackoverflow.com/questions/54858529/golang-reverse-a-arbitrary-slice
func Reverse(slice interface{}) {
	s := reflect.ValueOf(slice)

	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}

	swap := reflect.Swapper(s.Interface())
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
