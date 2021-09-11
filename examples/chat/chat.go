package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/librity/nc_nomadcoin/utils"
)

const (
	port    string = ":7000"
	baseURL string = "http://localhost" + port

	staticDir   string = "examples/chat"
	staticRoute string = "/static/"
)

var (
	handler *http.ServeMux = http.NewServeMux()

	upgrader = websocket.Upgrader{
		CheckOrigin: checkOrigin,
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
	handler.HandleFunc("/chat/ping", pingChat)
	handler.HandleFunc("/chat/echo", echoChat)
}

func listenOrDie() {
	fmt.Printf("📨 Chat server listening on %s\n", baseURL)
	fmt.Printf("➡️ Ping demo: %s%sping_chat.html\n", baseURL, staticRoute)
	fmt.Printf("➡️ Echo demo: %s%secho_chat.html\n", baseURL, staticRoute)

	log.Fatal(http.ListenAndServe(port, handler))
}

func pingChat(rw http.ResponseWriter, r *http.Request) {
	wsConn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleError(err)

	for {
		fmt.Println("Awaiting message...")
		_, payload, err := wsConn.ReadMessage()
		if err != nil {
			break
		}

		fmt.Printf("Message received: \"%s\"\n", payload)
		fmt.Println("---")
	}
}

func echoChat(rw http.ResponseWriter, r *http.Request) {
	wsConn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleError(err)

	for {
		fmt.Println("Awaiting message...")
		_, payload, err := wsConn.ReadMessage()
		if err != nil {
			break
		}

		fmt.Printf("Message received: \"%s\"\n", payload)
		time.Sleep(1 * time.Second)
		echo := fmt.Sprintf("ECHO: %s", payload)
		err = wsConn.WriteMessage(websocket.TextMessage, []byte(echo))
		if err != nil {
			break
		}

		fmt.Printf("Sending message \"%s\" to %s\n", echo, wsConn.RemoteAddr())
		fmt.Println("---")
	}
}

func checkOrigin(r *http.Request) bool { return true }
