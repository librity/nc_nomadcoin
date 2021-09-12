package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	port    = ":7000"
	baseURL = "http://localhost" + port

	staticDir   = "examples/chat"
	staticRoute = "/static/"
	staticURL   = baseURL + staticRoute
)

var (
	handler *http.ServeMux = http.NewServeMux()

	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
)

func main() {
	startServer()
}

func startServer() {
	loadFileServer()
	loadRoutes()
	listenOrDie()
}

func loadFileServer() {
	fileServer := http.FileServer(http.Dir(staticDir))
	handler.Handle(staticRoute, http.StripPrefix(staticRoute, fileServer))
}

func loadRoutes() {
	handler.HandleFunc("/chat/ping", pingDemo)
	handler.HandleFunc("/chat/echo", echoDemo)
	handler.HandleFunc("/chat/broadcast", broadcastDemo)
}

func listenOrDie() {
	fmt.Printf("📨 Chat server listening on %s\n", baseURL)
	fmt.Printf("➡️ Ping demo: %sping_demo.html\n", staticURL)
	fmt.Printf("➡️ Echo demo: %secho_demo.html\n", staticURL)
	fmt.Printf("➡️ Broadcast demo: %sbroadcast_demo.html\n", staticURL)
	fmt.Println("---")

	log.Fatal(http.ListenAndServe(port, handler))
}
