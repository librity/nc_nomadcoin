package blockchain

import (
	"fmt"
	"sync"

	"github.com/librity/nc_nomadcoin/utils"
)

type blockchain struct {
	LastHash  string
	Leight    int
	Dificulty int
	m         sync.Mutex
}

var (
	bc     *blockchain
	onceBC sync.Once
)

// Stringer interface: https://pkg.go.dev/fmt#Stringer
func (b *blockchain) String() string {
	s := fmt.Sprintln("=== Blockchain ===") +
		fmt.Sprintln("Height:", fmt.Sprint(b.Leight)) +
		fmt.Sprintln("Last Hash:", b.LastHash) +
		fmt.Sprintln("Current Dificulty:", b.Dificulty) +
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

	checkpoint := storage.LoadChain()
	if checkpoint == nil {
		bc.addBlock()
		fmt.Println("⛓️  Blockchain initialized succesfully.")
		return
	}

	bc.restore(checkpoint)
	fmt.Println("⛓️  Blockchain restored succesfully.")
}

func (b *blockchain) addBlock() *Block {
	block := createBlock(b.LastHash, b.Leight+1, getDifficulty(b))
	b.reference(block)

	return block
}

func (b *blockchain) reference(block *Block) {
	b.LastHash = block.Hash
	b.Leight = block.Height
	b.Dificulty = block.Difficulty
	b.save()
}

func (b *blockchain) save() {
	storage.SaveChain(utils.ToGob(b))
}

func (b *blockchain) restore(encoded []byte) {
	utils.FromGob(b, encoded)
}

func clearBC() {
	storage.ClearBlocks()
	storage.ClearChain()
}
