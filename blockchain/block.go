package blockchain

import (
	"errors"
	"fmt"
	"strings"

	"github.com/librity/nc_nomadcoin/db"
	"github.com/librity/nc_nomadcoin/utils"
)

var (
	ErrBlockNotFound = errors.New("block not found")
)

type Block struct {
	Height       int    `json:"height"`
	PreviousHash string `json:"previousHash,omitempty"`
	Hash         string `json:"hash"`
	Difficulty   int    `json:"difficulty"`
	NOnce        int    `json:"nOnce"`
	Timestamp    int64  `json:"timestamp"`
	Transactions []*Tx  `json:"transactions"`
}

func FindBlock(hash string) (*Block, error) {
	rawBlock := db.LoadBlock(hash)
	if rawBlock == nil {
		return nil, ErrBlockNotFound
	}

	block := blockFromBytes(rawBlock)
	return block, nil
}

func createBlock(prevHash string, height, difficulty int) *Block {
	block := newBlock(prevHash, height, difficulty)
	block.mine()
	block.loadTransactions()
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
	b.Transactions = Mempool.popAll()
}

func (b *Block) save() {
	db.SaveBlock(b.Hash, utils.ToGob(b))
}

func blockFromBytes(encoded []byte) *Block {
	block := &Block{}
	utils.FromGob(block, encoded)

	return block
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
	fmt.Println("Transactions:", b.Transactions)
	fmt.Println("---")
}
