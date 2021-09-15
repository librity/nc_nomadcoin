package db

import (
	"github.com/librity/nc_nomadcoin/utils"
	bolt "go.etcd.io/bbolt"
)

func SaveChain(chain []byte) {
	err := getDB().Update(func(transaction *bolt.Tx) error {
		bucket := transaction.Bucket([]byte(chainBucket))
		err := bucket.Put([]byte(chainCheckpoint), chain)

		return err
	})

	utils.PanicError(err)
}

func LoadChain() []byte {
	var chain []byte

	err := getDB().View(func(transaction *bolt.Tx) error {
		bucket := transaction.Bucket([]byte(chainBucket))
		chain = bucket.Get([]byte(chainCheckpoint))
		return nil
	})
	utils.PanicError(err)

	return chain
}

func ClearChain() {
	resetBucket(chainBucket)
}
