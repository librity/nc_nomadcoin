package main

import (
	"os"

	"github.com/librity/nc_nomadcoin/blockchain"
	"github.com/librity/nc_nomadcoin/cli"
	"github.com/librity/nc_nomadcoin/db"
	"github.com/librity/nc_nomadcoin/wallet"
)

func main() {
	startup()
	defer cleanExit()

	// go p2p.PingForever()

	// walletDemo()
	// printChainStatus()
	cli.Start()
}

func printChainStatus() {
	blockchain.InspectChain()
	blockchain.InspectBlocks()
}

func walletDemo() {
	wallet := wallet.GetW()
	wallet.Inspect()
}

func startup() {}

func cleanExit() {
	db.Close()
	os.Exit(0)
}
