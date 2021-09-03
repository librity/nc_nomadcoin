package main

import (
	"fmt"

	"github.com/librity/nc_nomadcoin/utils"
)

func main() {
	POWDemo()
}

func POWDemo() {
	transactions := "1 BTC from 6e13e92754d1 to 8fa99e8b244a,..."
	hash := utils.HexHash(transactions)

	fmt.Println(hash)
}
