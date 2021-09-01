package server

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port string = ":4000"
)

func Start() {
	loadTemplates()
	loadHandlers()
	listenOrDie()
}

func loadHandlers() {
	loadFileServer()
	loadRoutes()
}

func loadFileServer() {
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
}

func loadRoutes() {
	http.HandleFunc("/", home)
	http.HandleFunc("/blocks", blocks)
	http.HandleFunc("/blocks/new", newBlock)
}

func listenOrDie() {
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
