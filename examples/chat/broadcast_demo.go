package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/librity/nc_nomadcoin/utils"
)

var (
	activeConns []*websocket.Conn
)

const messageHTML = `
<p style="color: #FFFFFF; background-color: #4682B4;padding: 10px; border-radius: 5px;">
    From %s:
    <br/>
    %s
    <br/>
	%s
</p>
`

func broadcastDemo(rw http.ResponseWriter, r *http.Request) {
	thisConn, err := upgrader.Upgrade(rw, r, nil)
	utils.PanicError(err)

	thisAddress := thisConn.RemoteAddr()
	activeConns = append(activeConns, thisConn)
	for {
		fmt.Println(thisAddress, "awaiting message...")
		_, payload, err := thisConn.ReadMessage()
		if err != nil {
			break
		}

		timestamp := time.Now().Format(time.UnixDate)
		fmt.Printf("Message received: \"%s\" from %s\n", payload, thisAddress)
		message := fmt.Sprintf(messageHTML, thisAddress, payload, timestamp)
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

		fmt.Printf("Broadcasting to %s\n", conn.RemoteAddr())
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
