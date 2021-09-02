package cli

import (
	"flag"
	"os"

	"github.com/librity/nc_nomadcoin/explorer"
)

func handleExplorer() {
	var (
		command  = flag.NewFlagSet("explorer", flag.ExitOnError)
		portFlag = command.Int("port", 4000, "Sets the port of the server")
	)
	command.Parse(os.Args[2:])
	serverPort := *portFlag

	explorer.StartCustom(serverPort)
}
