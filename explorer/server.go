package explorer

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port        string = ":4000"
	staticDir   string = "explorer/static"
	staticRoute string = "/static/"
)

func Start() {
	loadTemplates()
	loadHandlers()
	listenOrDie()
}

func loadHandlers() {
	loadFileServer()
	loadRoutes()
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
	fmt.Printf("HTML Explorer listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
