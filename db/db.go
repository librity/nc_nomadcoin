package db

import (
	"fmt"
	"sync"

	"github.com/librity/nc_nomadcoin/utils"
	bolt "go.etcd.io/bbolt"
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

func Close() {
	if db == nil {
		return
	}

	getDB().Close()
	fmt.Println("🗃️  Database closed succesfully.")
}

func getDB() *bolt.DB {
	if db == nil {
		once.Do(initializeDB)
	}

	return db
}

func initializeDB() {
	openDB()
	createBuckets()

	fmt.Println("🗃️  Database initialized succesfully.")
}

func openDB() {
	dbPointer, err := bolt.Open(dbName, 0600, nil)
	utils.PanicError(err)
	db = dbPointer
}

func createBuckets() {
	err := db.Update(func(transaction *bolt.Tx) error {
		_, err := transaction.CreateBucketIfNotExists([]byte(chainBucket))
		utils.PanicError(err)
		_, err = transaction.CreateBucketIfNotExists([]byte(blocksBucket))

		return err
	})

	utils.PanicError(err)
}
