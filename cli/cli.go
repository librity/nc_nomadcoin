package cli

import (
	"os"
)

func Start() {
	checkArgs()
	handleCommand()
}

func checkArgs() {
	if len(os.Args) < 2 {
		printUsageAndDie()
	}
}

func handleCommand() {
	command := os.Args[1]

	switch command {
	case "explorer":
		handleExplorer()
	case "rest":
		handleRest()
	case "both":
		handleBoth()
	default:
		printUsageAndDie()
	}
}
