package db

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/librity/nc_nomadcoin/utils"
)

func SaveBlock(hash string, data []byte) {
	fmt.Printf("Block: %s\n", hash)
	fmt.Printf("Data: %b\n", data)

	err := Get().Update(func(transaction *bolt.Tx) error {
		bucket := transaction.Bucket([]byte(blocksBucket))
		err := bucket.Put([]byte(hash), data)

		return err
	})

	utils.HandleError(err)
}

func SaveBlockchain(chain []byte) {
	err := Get().Update(func(transaction *bolt.Tx) error {
		bucket := transaction.Bucket([]byte(chainBucket))
		err := bucket.Put([]byte("checkpoint"), chain)

		return err
	})

	utils.HandleError(err)
}
