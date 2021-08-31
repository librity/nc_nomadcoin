package main

import (
	"fmt"

	"github.com/librity/nc_nomadcoin/blockchain"
)

func main() {
	welcome()
	demo()
}

func welcome() {
	fmt.Println("Welcome to Nomad Coin!")
	fmt.Println("---")
}

func demo() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second block.")
	chain = blockchain.GetBlockchain()
	chain.AddBlock("Third block.")
	chain.AddBlock("4th block.")
	chain.AddBlock("5th block.")
	chain = blockchain.GetBlockchain()
	chain.AddBlock("6th block.")
	chain.ListBlocks()
}
