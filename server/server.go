package server

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const (
	port string = ":4000"
)

var templates *template.Template

func Start() {
	loadTemplates()
	loadHandlers()
	listenOrDie()
}

func loadHandlers() {
	http.HandleFunc("/", home)
}

func listenOrDie() {
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
