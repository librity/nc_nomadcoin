package explorer

import (
	"fmt"
	"log"
	"net/http"

	"github.com/librity/nc_nomadcoin/config"
	"github.com/librity/nc_nomadcoin/utils"
)

const (
	staticDir   string = "explorer/static"
	staticRoute string = "/static/"
)

var (
	port    string         = utils.BuildPort(config.DefaultExplorerPort)
	baseURL string         = "http://localhost" + port
	handler *http.ServeMux = http.NewServeMux()
)

func Start() {
	setEnvVars()

	loadTemplates()
	loadFileServer()
	loadRoutes()

	listenOrDie()
}

func setEnvVars() {
	portNum := config.GetRestPort()
	port = utils.BuildPort(portNum)
	baseURL = "http://localhost" + port
}

func loadFileServer() {
	fileServer := http.FileServer(http.Dir(staticDir))
	handler.Handle(staticRoute, http.StripPrefix(staticRoute, fileServer))
}

func loadRoutes() {
	handler.HandleFunc("/", home)
	handler.HandleFunc("/blocks", blocks)
	handler.HandleFunc("/blocks/new", newBlock)
}

func listenOrDie() {
	fmt.Printf("🧭 HTML Explorer listening on %s\n", baseURL)
	log.Fatal(http.ListenAndServe(port, handler))
}
