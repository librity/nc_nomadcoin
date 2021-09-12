package cli

import (
	"flag"
	"os"

	"github.com/librity/nc_nomadcoin/rest"
)

func handleRest() {
	var (
		command  = flag.NewFlagSet("rest", flag.ExitOnError)
		portFlag = command.Int("port", 5001, "Sets the port of the server")
	)
	command.Parse(os.Args[2:])
	serverPort := *portFlag

	rest.StartCustom(serverPort)
}
