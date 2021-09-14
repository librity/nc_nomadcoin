package cli

import (
	"flag"
	"os"

	"github.com/librity/nc_nomadcoin/config"
	"github.com/librity/nc_nomadcoin/explorer"
	"github.com/librity/nc_nomadcoin/rest"
)

func handleBoth() {
	var (
		command      = flag.NewFlagSet("both", flag.ExitOnError)
		explorerPort = command.Int("ePort", config.DefaultExplorerPort, "Sets the port of the HTML Explorer server")
		restPort     = command.Int("rPort", config.DefaultRestPort, "Sets the port of the REST API server")
	)

	command.Parse(os.Args[2:])
	config.SetExplorerPort(*explorerPort)
	config.SetRestPort(*restPort)

	go explorer.Start()
	rest.Start()
}
