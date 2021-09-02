package db

import (
	"github.com/boltdb/bolt"
	"github.com/librity/nc_nomadcoin/utils"
)

func SaveBlock(hash string, data []byte) {
	database := Get()
	err := database.Update(func(transaction *bolt.Tx) error {
		bucket := transaction.Bucket([]byte(blocksBucket))
		err := bucket.Put([]byte(hash), data)

		return err
	})

	utils.HandleError(err)
}
