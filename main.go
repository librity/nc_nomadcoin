package main

import (
	"os"

	"github.com/librity/nc_nomadcoin/blockchain"
	"github.com/librity/nc_nomadcoin/db"
	"github.com/librity/nc_nomadcoin/wallet"
)

func main() {
	defer cleanExit()

	walletDemo()
	// printChainStatus()
	// cli.Start()
}

func printChainStatus() {
	blockchain.InspectChain()
	blockchain.InspectBlocks()
}

func walletDemo() {
	wallet.Start()
}

func cleanExit() {
	db.Close()
	os.Exit(0)
}
