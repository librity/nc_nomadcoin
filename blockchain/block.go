package blockchain

import (
	"fmt"
	"strings"

	"github.com/librity/nc_nomadcoin/utils"
)

type Block struct {
	Height       int    `json:"height"`
	PreviousHash string `json:"previousHash,omitempty"`
	Hash         string `json:"hash"`
	Difficulty   int    `json:"difficulty"`
	NOnce        int    `json:"nOnce"`
	Timestamp    int64  `json:"timestamp"`
	Txs          []*Tx  `json:"transactions"`
}

func MineBlock() *Block {
	block := getBC().addBlock()

	return block
}

func (b *Block) mine() {
	target := strings.Repeat("0", b.Difficulty)

	for {
		b.Timestamp = now()
		attempt := utils.HexHash(b)

		if strings.HasPrefix(attempt, target) {
			b.Hash = attempt
			break
		}

		b.NOnce++
	}
}

func (b *Block) loadTransactions() {
	b.Txs = getMP().popAll()
}

func (b *Block) save() {
	storage.SaveBlock(b.Hash, utils.ToGob(b))
}

func (b *Block) inspect() {
	fmt.Println("Height:", b.Height)
	if b.PreviousHash != "" {
		fmt.Println("Previous hash:", b.PreviousHash)
	}
	fmt.Println("Hash:", b.Hash)
	fmt.Println("Difficulty:", b.Difficulty)
	fmt.Println("NOnce:", b.NOnce)
	fmt.Println("Timestamp:", b.Timestamp)
	fmt.Println("Transactions:", b.Txs)
	fmt.Println("---")
}

func createBlock(prevHash string, height, difficulty int) *Block {
	block := newBlock(prevHash, height, difficulty)
	block.loadTransactions()
	block.mine()
	block.save()

	return block
}

func newBlock(prevHash string, height, difficulty int) *Block {
	block := Block{
		Height:       height,
		PreviousHash: prevHash,
		Hash:         "",
		Difficulty:   difficulty,
		NOnce:        0,
	}

	return &block
}

func blockFromBytes(encoded []byte) *Block {
	block := &Block{}
	utils.FromGob(block, encoded)

	return block
}
