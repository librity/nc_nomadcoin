package cli

import (
	"fmt"
	"runtime"
)

const usage string = `Welcome to the Nomad Coin CLI!

Usage: go run main.go [commmad] [flags]

Examples:

	go run main.go rest -port=PORT			Start the REST API (recommended)
	go run main.go explorer -port=PORT		Start the HTML Explorer
	go run main.go both -ePort=PORT -rPort=PORT	Start both REST API and HTML Explorer

`

func printUsageAndDie() {
	fmt.Print(usage)
	runtime.Goexit()
}
