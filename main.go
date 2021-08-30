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
	chain.AddBlock("Genesis block.")
	chain.AddBlock("Second block.")
	chain = blockchain.GetBlockchain()
	chain.AddBlock("Third block.")
	chain.ListBlocks()
}
