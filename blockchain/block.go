package blockchain

import (
	"bytes"
	"encoding/gob"
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
}

func FindBlock(hash string) (*Block, error) {
	rawBlock := db.LoadBlock(hash)
	fmt.Println(rawBlock)
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

// Stringer interface: https://pkg.go.dev/fmt#Stringer
func (b Block) String() string {
	s := fmt.Sprintln("Height:", fmt.Sprint(b.Height)) +
		fmt.Sprintln("Data:", b.Data)
	if b.PreviousHash != "" {
		s = s + fmt.Sprintln("Previous hash:", b.PreviousHash)
	}
	s = s + fmt.Sprintln("Hash:", b.Hash) +
		fmt.Sprintln("---")

	return s
}

func (b *Block) listBlock() {
	fmt.Printf("Height: %d\n", b.Height)
	fmt.Printf("Data: %s\n", b.Data)
	if b.PreviousHash != "" {
		fmt.Printf("Previous hash: %s\n", b.PreviousHash)
	}
	fmt.Printf("Hash: %s\n", b.Hash)
	fmt.Println("---")
}

func blockFromBytes(encoded []byte) *Block {
	block := &Block{}
	buffer := bytes.NewReader(encoded)
	decoder := gob.NewDecoder(buffer)
	err := decoder.Decode(block)
	utils.HandleError(err)

	return block
}
