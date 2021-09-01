package rest

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port string = ":5000"
)

func Start() {
	loadHandlers()
	listenOrDie()
}

func loadHandlers() {
	http.HandleFunc("/", documentation)
}

func listenOrDie() {
	fmt.Printf("REST API listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
