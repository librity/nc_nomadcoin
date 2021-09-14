package config

import (
	"fmt"
	"sync"
)

const (
	DefaultDBName = "blockchain.db"
)

var (
	dbName     string
	dbNameOnce sync.Once
)

func GetDBName() string {
	if dbName == "" {
		return DefaultDBName
	}

	return dbName
}

func SetDBName() {
	dbNameOnce.Do(func() {
		dbName = buildDBName()
	})
}

func buildDBName() string {
	port := GetRestPort()
	dbName := fmt.Sprintf("blockchain_%d.db", port)

	return dbName
}
