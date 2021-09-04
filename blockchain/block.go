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
	Data         string `json:"data"`
	PreviousHash string `json:"previousHash,omitempty"`
	Hash         string `json:"hash"`
	Difficulty   int    `json:"difficulty"`
	NOnce        int    `json:"nOnce"`
	Timestamp    int    `json:"timestamp"`
}

func FindBlock(hash string) (*Block, error) {
	rawBlock := db.LoadBlock(hash)
	if rawBlock == nil {
		return nil, ErrBlockNotFound
	}

	block := blockFromBytes(rawBlock)
	return block, nil
}

func createBlock(data string, prevHash string, height int) *Block {
	block := newBlock(data, prevHash, height)
	block.mine()
	block.save()

	return block
}

func newBlock(data string, prevHash string, height int) *Block {
	block := Block{
		Height:       height,
		Data:         data,
		PreviousHash: prevHash,
		Hash:         "",
		Difficulty:   Get().difficulty(),
		NOnce:        0}

	return &block
}

func (b *Block) save() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

func (b *Block) mine() {
	target := strings.Repeat("0", b.Difficulty)

	for {
		b.Timestamp = utils.Now()
		attempt := utils.HexHash(b)

		if strings.HasPrefix(attempt, target) {
			b.Hash = attempt
			break
		}

		b.NOnce++
	}
}

func blockFromBytes(encoded []byte) *Block {
	block := &Block{}
	utils.FromBytes(block, encoded)

	return block
}

func (b *Block) inspect() {
	fmt.Println("Height:", b.Height)
	fmt.Println("Data:", b.Data)
	if b.PreviousHash != "" {
		fmt.Println("Previous hash:", b.PreviousHash)
	}
	fmt.Println("Hash:", b.Hash)
	fmt.Println("Difficulty:", b.Difficulty)
	fmt.Println("NOnce:", b.NOnce)
	fmt.Println("Timestamp:", b.Timestamp)
	fmt.Println("---")
}
