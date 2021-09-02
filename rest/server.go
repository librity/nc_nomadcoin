package rest

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port    string = ":5000"
	baseURL string = "http://localhost" + port
)

func Start() {
	loadHandlers()
	listenOrDie()
}

func loadHandlers() {
	http.HandleFunc("/", documentation)
	http.HandleFunc("/blocks", blocks)
}

func listenOrDie() {
	fmt.Printf("REST API listening on %s\n", baseURL)
	log.Fatal(http.ListenAndServe(port, nil))
}
