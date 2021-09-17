package blockchain

import "github.com/librity/nc_nomadcoin/db"

type storageLayerI interface {
	LoadChain() []byte
	SaveChain(chain []byte)
	ClearChain()

	LoadBlock(hash string) []byte
	SaveBlock(hash string, data []byte)
	ClearBlocks()
}

var storage storageLayerI = db.DBLayer{}
