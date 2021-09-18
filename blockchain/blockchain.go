package blockchain

import (
	"fmt"
	"sync"

	"github.com/librity/nc_nomadcoin/utils"
)

type blockchain struct {
	LastHash   string
	Height     int
	Difficulty int
	m          sync.Mutex
}

var (
	bc     *blockchain
	onceBC sync.Once
)

// Stringer interface: https://pkg.go.dev/fmt#Stringer
func (b *blockchain) String() string {
	s := fmt.Sprintln("=== Blockchain ===") +
		fmt.Sprintln("Height:", fmt.Sprint(b.Height)) +
		fmt.Sprintln("Last Hash:", b.LastHash) +
		fmt.Sprintln("Current Dificulty:", b.Difficulty) +
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
	block := createBlock(b.LastHash, b.Height+1, getDifficulty(b))
	b.reference(block)

	return block
}

func (b *blockchain) reference(block *Block) {
	b.LastHash = block.Hash
	b.Height = block.Height
	b.Difficulty = block.Difficulty
	b.save()
}

func (b *blockchain) save() {
	storage.SaveChain(utils.ToGob(b))
}

func (b *blockchain) restore(encoded []byte) {
	utils.FromGob(b, encoded)
}

func (b *blockchain) reset() {
	storage.ClearBlocks()
	storage.ClearChain()
}

func newBC(lastHash string, height, difficulty int) *blockchain {
	chain := &blockchain{
		LastHash:   lastHash,
		Height:     height,
		Difficulty: difficulty,
	}

	return chain
}
