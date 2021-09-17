package db

import (
	"github.com/librity/nc_nomadcoin/utils"
	bolt "go.etcd.io/bbolt"
)

func saveChain(chain []byte) {
	err := getDB().Update(func(transaction *bolt.Tx) error {
		bucket := transaction.Bucket([]byte(chainBucket))
		err := bucket.Put([]byte(chainKey), chain)

		return err
	})

	utils.PanicError(err)
}

func loadChain() []byte {
	var chain []byte

	err := getDB().View(func(transaction *bolt.Tx) error {
		bucket := transaction.Bucket([]byte(chainBucket))
		chain = bucket.Get([]byte(chainKey))
		return nil
	})
	utils.PanicError(err)

	return chain
}

func clearChain() {
	resetBucket(chainBucket)
}
