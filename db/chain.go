package db

import (
	"github.com/boltdb/bolt"
	"github.com/librity/nc_nomadcoin/utils"
)

func SaveCheckpoint(chain []byte) {
	database := Get()
	err := database.Update(func(transaction *bolt.Tx) error {
		bucket := transaction.Bucket([]byte(chainBucket))
		err := bucket.Put([]byte(chainCheckpoint), chain)

		return err
	})

	utils.HandleError(err)
}

func LoadCheckpoint() []byte {
	var chain []byte

	database := Get()
	err := database.View(func(transaction *bolt.Tx) error {
		bucket := transaction.Bucket([]byte(chainBucket))
		chain = bucket.Get([]byte(chainCheckpoint))
		return nil
	})
	utils.HandleError(err)

	return chain
}
