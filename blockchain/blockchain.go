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
	m                sync.Mutex
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
		fmt.Sprintln("Current Dificulty:", b.CurrentDificulty) +
		fmt.Sprintln("")

	return s
}

func InspectChain() {
	chain := getBC()
	fmt.Print(chain)
}

func Status() blockchain {
	chain := getBC()

	return *chain
}

func getBC() *blockchain {
	once.Do(initializeBlockchain)

	return b
}

func initializeBlockchain() {
	b = &blockchain{Height: 0}
	b.m.Lock()
	defer b.m.Unlock()

	checkpoint := db.LoadChain()
	if checkpoint == nil {
		b.addBlock()
		fmt.Println("⛓️  Blockchain initialized succesfully.")
		return
	}

	b.restore(checkpoint)
	fmt.Println("⛓️  Blockchain restored succesfully.")
}

func (b *blockchain) addBlock() *Block {
	block := createBlock(b.LastHash, b.Height+1, getDifficulty(b))
	b.reference(block)

	return block
}

func (b *blockchain) reference(block *Block) {
	b.LastHash = block.Hash
	b.Height = block.Height
	b.CurrentDificulty = block.Difficulty
	b.save()
}

func (b *blockchain) save() {
	db.SaveChain(utils.ToGob(b))
}

func (b *blockchain) restore(encoded []byte) {
	utils.FromGob(b, encoded)
}

func (b *blockchain) reset() {
	db.ClearBlocks()
	db.ClearChain()
}
