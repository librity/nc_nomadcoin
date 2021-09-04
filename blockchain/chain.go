package blockchain

import (
	"fmt"
	"sync"

	"github.com/librity/nc_nomadcoin/db"
	"github.com/librity/nc_nomadcoin/utils"
)

type blockchain struct {
	LastHash         string `json:"lastHash"`
	Height           int    `json:"height"`
	CurrentDificulty int    `json:"currentDifficulty"`
}

var (
	b    *blockchain
	once sync.Once
)

// Stringer interface: https://pkg.go.dev/fmt#Stringer
func (b blockchain) String() string {
	s := fmt.Sprintln("=== Blockchain ===") +
		fmt.Sprintln("Height:", fmt.Sprint(b.Height)) +
		fmt.Sprintln("Last Hash:", b.LastHash) +
		fmt.Sprintln("")

	return s
}

func Get() *blockchain {
	if b == nil {
		once.Do(initializeBlockchain)
	}

	return b
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.LastHash, b.Height+1)
	b.reference(block)
}

func initializeBlockchain() {
	b = &blockchain{Height: 0}

	checkpoint := db.LoadCheckpoint()
	if checkpoint == nil {
		b.AddBlock("The Times 03/Jan/2009 Chancellor on brink of second bailout for banks")
		return
	}

	b.restore(checkpoint)
}

func (b *blockchain) reference(block *Block) {
	b.LastHash = block.Hash
	b.Height = block.Height
	b.CurrentDificulty = block.Difficulty
	b.save()
}

func (b *blockchain) save() {
	db.SaveCheckpoint(utils.ToBytes(b))
}

func (b *blockchain) restore(encoded []byte) {
	utils.FromBytes(b, encoded)
}
