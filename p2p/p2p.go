package p2p

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/librity/nc_nomadcoin/utils"
)

func AddPeer(address, port, thisPort string) {
	format := "ws://%s:%s/peers/upgrade?thisPort=%s"
	url := fmt.Sprintf(format, address, port, thisPort)
	juniorConn, _, err := websocket.DefaultDialer.Dial(url, nil)
	utils.PanicError(err)

	initPeer(address, port, juniorConn)
}
