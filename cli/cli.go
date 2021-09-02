package cli

import (
	"fmt"
	"os"
)

func Start() {
	if len(os.Args) < 2 {
		printUsageAndDie()
	}

	command := os.Args[1]
	switch command {
	case "explorer":
		fmt.Println("Start Explorer")
	case "rest":
		fmt.Println("Start REST API")
	default:
		printUsageAndDie()
	}
}

const usage string = `Welcome to the Nomad Coin CLI!

Please use one of the following commands:

explorer PORT	Start the HTLM Explorer
rest PORT		Start the REST API (recommended)

---
`

func printUsageAndDie() {
	fmt.Print(usage)
	os.Exit(0)
}
