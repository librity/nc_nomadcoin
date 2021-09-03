package db

import (
	"github.com/boltdb/bolt"
	"github.com/librity/nc_nomadcoin/utils"
)

func SaveBlock(hash string, data []byte) {
	err := getDB().Update(func(transaction *bolt.Tx) error {
		bucket := transaction.Bucket([]byte(blocksBucket))
		err := bucket.Put([]byte(hash), data)

		return err
	})

	utils.HandleError(err)
}

func LoadBlock(hash string) []byte {
	var rawBlock []byte

	err := getDB().View(func(transaction *bolt.Tx) error {
		bucket := transaction.Bucket([]byte(blocksBucket))
		rawBlock = bucket.Get([]byte(hash))
		return nil
	})
	utils.HandleError(err)

	return rawBlock
}
