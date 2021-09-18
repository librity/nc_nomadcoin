package explorer

import (
	"fmt"
	"log"
	"net/http"

	"github.com/librity/nc_nomadcoin/config"
	"github.com/librity/nc_nomadcoin/utils"
)

var (
	port    string = utils.BuildPort(config.DefaultExplorerPort)
	baseURL string = "http://localhost" + port
)

func Start() {
	setEnvVars()

	loadTemplates()
	loadFileServer()
	loadRoutes()

	listenOrDie()
}

func setEnvVars() {
	portNum := config.GetExplorerPort()
	port = utils.BuildPort(portNum)
	baseURL = "http://localhost" + port
}

func listenOrDie() {
	fmt.Printf("🧭 HTML Explorer listening on %s\n", baseURL)
	log.Fatal(http.ListenAndServe(port, router))
}
