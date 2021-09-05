package main

import (
	"fmt"

	"github.com/librity/nc_nomadcoin/blockchain"
	"github.com/librity/nc_nomadcoin/cli"
	"github.com/librity/nc_nomadcoin/db"
)

func main() {
	defer db.Close()

	blockchainDemo()
	cliDemo()
}

func blockchainDemo() {
	chain := blockchain.Get()
	fmt.Print(chain)

	// seedBlocks()

	chain.ListBlocks()
}

func seedBlocks() {
	chain := blockchain.Get()
	chain.AddBlock()
	chain.AddBlock()
}

func cliDemo() {
	cli.Start()
}
