package db

import (
	"sync"

	"github.com/boltdb/bolt"
	"github.com/librity/nc_nomadcoin/utils"
)

const (
	dbName = "blockchain.db"

	chainBucket     = "chain"
	chainCheckpoint = "checkpoint"

	blocksBucket = "blocks"
)

var (
	db   *bolt.DB
	once sync.Once
)

func getDB() *bolt.DB {
	if db == nil {
		once.Do(initializeDB)
	}

	return db
}

func initializeDB() {
	openDB()
	createBuckets()
}

func openDB() {
	dbPointer, err := bolt.Open(dbName, 0600, nil)
	utils.HandleError(err)
	db = dbPointer
}

func createBuckets() {
	err := db.Update(func(transaction *bolt.Tx) error {
		_, err := transaction.CreateBucketIfNotExists([]byte(chainBucket))
		utils.HandleError(err)
		_, err = transaction.CreateBucketIfNotExists([]byte(blocksBucket))

		return err
	})

	utils.HandleError(err)
}
