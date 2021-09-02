package rest

import (
	"fmt"
	"log"
	"net/http"
)

var (
	port    string = ":5000"
	baseURL string = "http://localhost" + port
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
	listenOrDie()
}

func setEnvVars(portNum int) {
	port = fmt.Sprintf(":%d", portNum)
	baseURL = "http://localhost" + port
}

func loadHandlers() {
	http.HandleFunc("/", documentation)
	http.HandleFunc("/blocks", blocks)
}

func listenOrDie() {
	fmt.Printf("REST API listening on %s\n", baseURL)
	log.Fatal(http.ListenAndServe(port, nil))
}
