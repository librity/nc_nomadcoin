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
	loadRoutes()
	loadMiddlewares()
	listenOrDie()
}

func setEnvVars(portNum int) {
	port = fmt.Sprintf(":%d", portNum)
	baseURL = "http://localhost" + port
}

func listenOrDie() {
	fmt.Printf("🤖 REST API listening on %s\n", baseURL)
	log.Fatal(http.ListenAndServe(port, router))
}
