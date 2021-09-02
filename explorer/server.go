package explorer

import (
	"fmt"
	"log"
	"net/http"
)

const (
	staticDir   string = "explorer/static"
	staticRoute string = "/static/"
)

var (
	port    string = ":4000"
	baseURL string = "http://localhost" + port
)

func Start() {
	loadAndListen()
}

func StartCustom(portNum int) {
	setEnvVars(portNum)
	loadAndListen()
}

func setEnvVars(portNum int) {
	port = fmt.Sprintf(":%d", portNum)
	baseURL = "http://localhost" + port
}

func loadAndListen() {
	loadTemplates()
	loadFileServer()
	loadRoutes()

	listenOrDie()
}

func loadFileServer() {
	fileServer := http.FileServer(http.Dir(staticDir))
	http.Handle(staticRoute, http.StripPrefix(staticRoute, fileServer))
}

func loadRoutes() {
	http.HandleFunc("/", home)
	http.HandleFunc("/blocks", blocks)
	http.HandleFunc("/blocks/new", newBlock)
}

func listenOrDie() {
	fmt.Printf("HTML Explorer listening on %s\n", baseURL)
	log.Fatal(http.ListenAndServe(port, nil))
}
