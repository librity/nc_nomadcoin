package cli

import (
	"flag"
	"os"

	"github.com/librity/nc_nomadcoin/explorer"
	"github.com/librity/nc_nomadcoin/rest"
)

func handleBoth() {
	var (
		command   = flag.NewFlagSet("both", flag.ExitOnError)
		ePortFlag = command.Int("ePort", 4000, "Sets the port of the HTML Explorer server")
		rPortFlag = command.Int("rPort", 5001, "Sets the port of the REST API server")
	)
	command.Parse(os.Args[2:])
	ePort := *ePortFlag
	rPort := *rPortFlag

	go explorer.StartCustom(ePort)
	rest.StartCustom(rPort)
}
