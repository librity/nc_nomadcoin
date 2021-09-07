package main

import (
	"fmt"

	"github.com/librity/nc_nomadcoin/blockchain"
	"github.com/librity/nc_nomadcoin/cli"
	"github.com/librity/nc_nomadcoin/db"
)

func main() {
	defer db.Close()

	printChainStatus()
	cli.Start()
}

func printChainStatus() {
	chain := blockchain.Get()
	fmt.Print(chain)

	chain.ListBlocks()
}
