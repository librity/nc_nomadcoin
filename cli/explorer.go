package cli

import (
	"flag"
	"os"

	"github.com/librity/nc_nomadcoin/config"
	"github.com/librity/nc_nomadcoin/explorer"
)

func handleExplorer() {
	var (
		command      = flag.NewFlagSet("explorer", flag.ExitOnError)
		explorerPort = command.Int("port", config.DefaultExplorerPort, "Sets the port of the server")
	)

	command.Parse(os.Args[2:])
	config.SetExplorerPort(*explorerPort)

	explorer.Start()
}
