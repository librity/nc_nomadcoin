package p2p

import (
	"fmt"
	"net/http"
	"time"

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
	}
}

// NOT SAFE AT ALL
func checkOrigin(r *http.Request) bool { return true }
