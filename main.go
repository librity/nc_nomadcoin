package main

import (
	"github.com/librity/nc_nomadcoin/blockchain"
	"github.com/librity/nc_nomadcoin/db"
	"github.com/librity/nc_nomadcoin/wallet"
)

func main() {
	defer db.Close()

	// printChainStatus()
	// cli.Start()
	walletDemo()
}

func printChainStatus() {
	blockchain.InspectChain()
	blockchain.InspectBlocks()
}

func walletDemo() {
	wallet.Start()
}
