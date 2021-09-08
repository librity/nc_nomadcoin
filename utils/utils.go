package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"
	"reflect"
)

var (
	ErrBigIntBadBytes = errors.New("bytes array should be 32 in length")
)

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

func HandleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func HexHash(data interface{}) string {
	defaultFormat := fmt.Sprintf("%v", data)

	return HexHashStr(defaultFormat)
}

func HexHashStr(data string) string {
	dataBytes := []byte(data)
	rawHash := sha256.Sum256(dataBytes)
	hexHash := fmt.Sprintf("%x", rawHash)

	return hexHash
}

func HexToBytes(hexStr string) []byte {
	bytes, err := hex.DecodeString(hexStr)
	HandleError(err)

	return bytes
}

func BytesToHex(bytes []byte) string {
	hexStr := hex.EncodeToString(bytes)

	return hexStr
}

func BytesToBigInt(bytes []byte) (*big.Int, error) {
	if len(bytes) != 32 {
		return nil, ErrBigIntBadBytes
	}

	bi := new(big.Int).SetBytes(bytes)
	return bi, nil
}

func ToGob(i interface{}) []byte {
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(i)
	HandleError(err)

	return buffer.Bytes()
}

func FromGob(target interface{}, encoded []byte) {
	buffer := bytes.NewReader(encoded)
	decoder := gob.NewDecoder(buffer)
	err := decoder.Decode(target)
	HandleError(err)
}
