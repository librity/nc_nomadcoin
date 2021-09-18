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
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/404", notFound).Methods("GET")

	router.HandleFunc("/blocks", blocksIndex).Methods("GET")
	router.HandleFunc("/blocks/{hash:[0-9a-f]+}", blocksShow).Methods("GET")
	router.HandleFunc("/blocks", blocksCreate).Methods("POST")
	router.HandleFunc("/blocks/mine", blocksMine).Methods("GET")
}
