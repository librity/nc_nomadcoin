package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	port    string      = ":5000"
	baseURL string      = "http://localhost" + port
	router  *mux.Router = mux.NewRouter()
)

func Start() {
	loadAndListen()
}

func StartCustom(portNum int) {
	setEnvVars(portNum)
	loadAndListen()
}

func loadAndListen() {
	loadHandlers()
	loadMiddlewares()
	listenOrDie()
}

func setEnvVars(portNum int) {
	port = fmt.Sprintf(":%d", portNum)
	baseURL = "http://localhost" + port
}

func loadHandlers() {
	router.HandleFunc("/", documentation).Methods("GET")

	// router.HandleFunc("/blocks", blocksIndex).Methods("GET")
	router.HandleFunc("/blocks", createBlock).Methods("POST")
	router.HandleFunc("/blocks/{hash:[0-9a-f]+}", block).Methods("GET")
}

func loadMiddlewares() {
	router.Use(jsonContentTypeMiddleware)
}

func listenOrDie() {
	fmt.Printf("REST API listening on %s\n", baseURL)
	log.Fatal(http.ListenAndServe(port, router))
}
