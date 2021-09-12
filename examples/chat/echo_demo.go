package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/librity/nc_nomadcoin/utils"
)

func echoDemo(rw http.ResponseWriter, r *http.Request) {
	wsConn, err := upgrader.Upgrade(rw, r, nil)
	utils.PanicError(err)

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
