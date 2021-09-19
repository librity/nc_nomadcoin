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

	router.HandleFunc("/transactions/{id:[0-9a-f]+}", txsShow).Methods("GET")

	router.HandleFunc("/wallets", walletsIndex).Methods("GET")
	router.HandleFunc("/wallets/server", walletsServer).Methods("GET")
	router.HandleFunc("/wallets/{address:[0-9a-f]+}", walletsShow).Methods("GET")
}
