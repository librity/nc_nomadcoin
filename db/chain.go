package db

import (
	"github.com/boltdb/bolt"
	"github.com/librity/nc_nomadcoin/utils"
)

func SaveChain(chain []byte) {
	database := Get()
	err := database.Update(func(transaction *bolt.Tx) error {
		bucket := transaction.Bucket([]byte(chainBucket))
		err := bucket.Put([]byte("checkpoint"), chain)

		return err
	})

	utils.HandleError(err)
}
