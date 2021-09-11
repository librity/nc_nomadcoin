package p2p

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/librity/nc_nomadcoin/utils"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: checkOrigin,
	}
)

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	wsConn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleError(err)

	for {
		fmt.Println("Awaiting message...")
		_, payload, err := wsConn.ReadMessage()
		fmt.Println("Message received!")
		utils.HandleError(err)
		fmt.Printf("\"%s\"\n", payload)
	}
}

// NOT SAFE AT ALL
func checkOrigin(r *http.Request) bool { return true }
