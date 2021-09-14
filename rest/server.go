package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/librity/nc_nomadcoin/config"
	"github.com/librity/nc_nomadcoin/utils"
)

var (
	port    string = utils.BuildPort(config.DefaultRestPort)
	baseURL string = "http://localhost" + port
)

func Start() {
	setEnvVars()

	loadRoutes()
	loadMiddlewares()

	listenOrDie()
}

func setEnvVars() {
	portNum := config.GetRestPort()
	port = utils.BuildPort(portNum)
	baseURL = "http://localhost" + port
}

func listenOrDie() {
	fmt.Printf("🤖 REST API listening on %s\n", baseURL)
	log.Fatal(http.ListenAndServe(port, router))
}
