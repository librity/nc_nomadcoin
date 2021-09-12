package p2p

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/librity/nc_nomadcoin/utils"
)

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	seniorConn, err := upgrader.Upgrade(rw, r, nil)
	utils.PanicError(err)

}

func AddPeer(address, port string) {
	url := fmt.Sprintf("ws://%s:%s/peers/upgrade", address, port)
	juniorConn, _, err := websocket.DefaultDialer.Dial(url, nil)
	utils.PanicError(err)

}
