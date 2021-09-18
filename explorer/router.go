package explorer

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	staticDir   string = "explorer/static"
	staticRoute string = "/static/"
)

var (
	router *mux.Router = mux.NewRouter()
)

func loadFileServer() {
	fileServer := http.FileServer(http.Dir(staticDir))
	router.Handle(staticRoute, http.StripPrefix(staticRoute, fileServer))
}

func loadRoutes() {
	router.HandleFunc("/", home)

	router.HandleFunc("/blocks", blocksIndex).Methods("GET")
	router.HandleFunc("/blocks", createBlock).Methods("POST")
	router.HandleFunc("/blocks/mine", mineBlock).Methods("GET")
}
