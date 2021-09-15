package db

import (
	"github.com/librity/nc_nomadcoin/utils"
	bolt "go.etcd.io/bbolt"
)

func emptyBucket(name string) {
	err := getDB().Update(func(transaction *bolt.Tx) error {
		bucket := transaction.Bucket([]byte(name))
		err := bucket.ForEach(func(key, value []byte) error {
			err := bucket.Delete(key)

			return err
		})

		return err
	})

	utils.PanicError(err)
}

func resetBucket(name string) {
	bytesName := []byte(name)
	err := getDB().Update(func(transaction *bolt.Tx) error {
		err := transaction.DeleteBucket(bytesName)
		utils.PanicError(err)

		_, err = transaction.CreateBucketIfNotExists(bytesName)
		return err
	})

	utils.PanicError(err)
}
