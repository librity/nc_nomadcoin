package p2p

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/librity/nc_nomadcoin/utils"
)

const (
	WSURLformat = "ws://%s:%s/peers/upgrade?thisPort=%s"
)

func AddPeer(ip, port, thisPort string) {
	url := makeWSURL(ip, port, thisPort)
	seniorConn, _, err := websocket.DefaultDialer.Dial(url, nil)
	utils.PanicError(err)

	peer := initPeer(ip, port, seniorConn)
	peer.inbox <- []byte("I am junior, you are senior.")
}

func makeWSURL(ip, port, thisPort string) string {
	url := fmt.Sprintf(WSURLformat, ip, port, thisPort)

	return url
}
