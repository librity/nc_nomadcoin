package p2p

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: checkOrigin,
	}
)

// NOT SAFE AT ALL
func checkOrigin(r *http.Request) bool { return true }
