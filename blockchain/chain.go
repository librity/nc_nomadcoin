package blockchain

import (
	"fmt"
	"sync"

	"github.com/librity/nc_nomadcoin/db"
	"github.com/librity/nc_nomadcoin/utils"
)

type blockchain struct {
	LastHash string `json:"newestHash"`
	Height   int    `json:"height"`
}

var (
	b    *blockchain
	once sync.Once
)

func Get() *blockchain {
	if b == nil {
		once.Do(initializeBlockchain)
	}

	return b
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.LastHash, b.Height+1)
	b.mine(block)
}

func (b *blockchain) ListBlocks() {
}

func (b *blockchain) AllBlocks() {
}

func initializeBlockchain() {
	b = &blockchain{"", 0}

	checkpoint := db.LoadCheckpoint()
	if checkpoint == nil {
		b.AddBlock("The Times 03/Jan/2009 Chancellor on brink of second bailout for banks")
		return
	}

	b.restore(checkpoint)
}

func (b *blockchain) mine(block *Block) {
	b.LastHash = block.Hash
	b.Height = block.Height
	b.save()
}

func (b *blockchain) save() {
	db.SaveCheckpoint(utils.ToBytes(b))
}

func (b *blockchain) restore(encoded []byte) {
	utils.FromBytes(b, encoded)
}

func (b blockchain) String() string {
	s := fmt.Sprintln("=== Blockchain ===") +
		fmt.Sprintln("Height:", fmt.Sprint(b.Height)) +
		fmt.Sprintln("Last Hash:", b.LastHash) +
		"---"

	return s
}
