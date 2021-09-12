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
	activeConns []*websocket.Conn
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

func pingDemo(rw http.ResponseWriter, r *http.Request) {
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

func echoDemo(rw http.ResponseWriter, r *http.Request) {
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

func broadcastDemo(rw http.ResponseWriter, r *http.Request) {
	thisConn, err := upgrader.Upgrade(rw, r, nil)
	thisAddress := thisConn.RemoteAddr()
	utils.HandleError(err)

	activeConns = append(activeConns, thisConn)
	for {
		fmt.Println(thisAddress, "awaiting message...")
		_, payload, err := thisConn.ReadMessage()
		if err != nil {
			break
		}

		fmt.Printf("Message received: \"%s\" from %s\n", payload, thisAddress)
		message := fmt.Sprintf("From %s: %s", thisAddress, payload)
		badConns := broadcastMessage(message)
		cleanupConns(badConns)

		fmt.Println("---")
	}
}

func broadcastMessage(message string) []*websocket.Conn {
	var badConns []*websocket.Conn

	for _, conn := range activeConns {
		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			badConns = append(badConns, conn)
			continue
		}

		fmt.Printf("Broadcasting \"%s\" to %s\n", message, conn.RemoteAddr())
	}

	return badConns
}

func cleanupConns(badConns []*websocket.Conn) {
	for _, badConn := range badConns {
		fmt.Println("Removing", badConn.RemoteAddr(), "from broadcast.")
		removeConn(badConn)
	}
}

func removeConn(target *websocket.Conn) {
	index := -1
	for i, conn := range activeConns {
		if conn == target {
			index = i
			break
		}
	}

	if index == -1 {
		return
	}

	removeConnByIndex(index)
}

func removeConnByIndex(index int) {
	activeConns = append(activeConns[:index], activeConns[index+1:]...)
}
