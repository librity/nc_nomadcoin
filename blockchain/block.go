package blockchain

import (
	"errors"
	"fmt"

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
	block.save()

	return block
}

func newBlock(data string, prevHash string, height int) *Block {
	block := Block{
		Height:       height,
		Data:         data,
		PreviousHash: prevHash,
		Hash:         ""}
	block.setHash()

	return &block
}

func (b *Block) save() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

func (b *Block) setHash() {
	b.Hash = b.generateHash()
}

func (b *Block) generateHash() string {
	signature := b.Data + b.PreviousHash + fmt.Sprint(b.Height)

	return utils.HexHash(signature)
}

func blockFromBytes(encoded []byte) *Block {
	block := &Block{}
	utils.FromBytes(block, encoded)

	return block
}

// Stringer interface: https://pkg.go.dev/fmt#Stringer
func (b Block) String() string {
	s := fmt.Sprintln("Height:", b.Height) +
		fmt.Sprintln("Data:", b.Data)
	if b.PreviousHash != "" {
		s = s + fmt.Sprintln("Previous hash:", b.PreviousHash)
	}
	s = s + fmt.Sprintln("Hash:", b.Hash) +
		fmt.Sprintln("Difficulty:", b.Difficulty) +
		fmt.Sprintln("NOnce:", b.NOnce) +
		fmt.Sprintln("---")

	return s
}
