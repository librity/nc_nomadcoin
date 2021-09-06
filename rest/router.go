package rest

import "github.com/gorilla/mux"

var (
	router *mux.Router = mux.NewRouter()
)

func loadRoutes() {
	router.HandleFunc("/", documentation).Methods("GET")

	router.HandleFunc("/blokchain", blokchain).Methods("GET")

	router.HandleFunc("/blocks", blocksIndex).Methods("GET")
	router.HandleFunc("/blocks", createBlock).Methods("POST")
	router.HandleFunc("/blocks/{hash:[0-9a-f]+}", block).Methods("GET")

	router.HandleFunc("/wallet/{address}", wallet).Methods("GET")

	router.HandleFunc("/mempool", mempool).Methods("GET")

	router.HandleFunc("/transactions", createTransaction).Methods("POST")
}

func loadMiddlewares() {
	router.Use(jsonContentTypeMiddleware)
}
