package cli

import (
	"flag"
	"os"

	"github.com/librity/nc_nomadcoin/config"
	"github.com/librity/nc_nomadcoin/rest"
)

func handleRest() {
	var (
		command  = flag.NewFlagSet("rest", flag.ExitOnError)
		restPort = command.Int("port", config.DefaultRestPort, "Sets the port of the server")
	)

	command.Parse(os.Args[2:])
	config.SetRestPort(*restPort)

	rest.Start()
}
