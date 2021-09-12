package db

import (
	"github.com/boltdb/bolt"
	"github.com/librity/nc_nomadcoin/utils"
)

func SaveCheckpoint(chain []byte) {
	err := getDB().Update(func(transaction *bolt.Tx) error {
		bucket := transaction.Bucket([]byte(chainBucket))
		err := bucket.Put([]byte(chainCheckpoint), chain)

		return err
	})

	utils.PanicError(err)
}

func LoadCheckpoint() []byte {
	var chain []byte

	err := getDB().View(func(transaction *bolt.Tx) error {
		bucket := transaction.Bucket([]byte(chainBucket))
		chain = bucket.Get([]byte(chainCheckpoint))
		return nil
	})
	utils.PanicError(err)

	return chain
}
