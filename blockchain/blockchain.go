package blockchain

import (
	"fmt"
	"sync"

	"github.com/librity/nc_nomadcoin/db"
	"github.com/librity/nc_nomadcoin/utils"
)

type blockchain struct {
	lastHash  string
	height    int
	dificulty int
	m         sync.Mutex
}

var (
	bc     *blockchain
	onceBC sync.Once
)

// Stringer interface: https://pkg.go.dev/fmt#Stringer
func (b *blockchain) String() string {
	s := fmt.Sprintln("=== Blockchain ===") +
		fmt.Sprintln("Height:", fmt.Sprint(b.height)) +
		fmt.Sprintln("Last Hash:", b.lastHash) +
		fmt.Sprintln("Current Dificulty:", b.dificulty) +
		fmt.Sprintln("")

	return s
}

func InspectChain() {
	chain := getBC()
	fmt.Print(chain)
}

func getBC() *blockchain {
	onceBC.Do(initializeBC)

	return bc
}

func initializeBC() {
	bc = &blockchain{}
	bc.m.Lock()
	defer bc.m.Unlock()

	checkpoint := db.LoadChain()
	if checkpoint == nil {
		bc.addBlock()
		fmt.Println("⛓️  Blockchain initialized succesfully.")
		return
	}

	bc.restore(checkpoint)
	fmt.Println("⛓️  Blockchain restored succesfully.")
}

func (b *blockchain) addBlock() *Block {
	block := createBlock(b.lastHash, b.height+1, getDifficulty(b))
	b.reference(block)

	return block
}

func (b *blockchain) reference(block *Block) {
	b.lastHash = block.Hash
	b.height = block.Height
	b.dificulty = block.Difficulty
	b.save()
}

func (b *blockchain) save() {
	db.SaveChain(utils.ToGob(b))
}

func (b *blockchain) restore(encoded []byte) {
	utils.FromGob(b, encoded)
}

func resetBC() {
	db.ClearBlocks()
	db.ClearChain()
}
