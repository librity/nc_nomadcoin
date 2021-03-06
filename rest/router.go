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

	router.HandleFunc("/wallets", wltIndex).Methods("GET")
	router.HandleFunc("/wallets/server", serverWlt).Methods("GET")
	router.HandleFunc("/wallets/{address:[0-9a-f]+}", wlt).Methods("GET")

	router.HandleFunc("/mempool", mempool).Methods("GET")

	router.HandleFunc("/transactions", txIndex).Methods("GET")
	router.HandleFunc("/transactions", createTx).Methods("POST")

	router.HandleFunc("/peers", peersIndex).Methods("GET")
	router.HandleFunc("/peers", addPeer).Methods("POST")
	router.HandleFunc("/peers/upgrade", upgradePeer).Methods("GET")
}

func loadMiddlewares() {
	router.Use(loggerMw)
	router.Use(jsonContentTypeMw)
}
