package p2p

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/librity/nc_nomadcoin/config"
	"github.com/librity/nc_nomadcoin/utils"
)

const (
	WSURLformat = "ws://%s:%s/peers/upgrade?thisPort=%s"
)

func AddPeer(ip, port string) {
	url := makeWSURL(ip, port)
	seniorConn, _, err := websocket.DefaultDialer.Dial(url, nil)
	utils.PanicError(err)

	peer := initPeer(ip, port, seniorConn)
	go broadcastNewPeer(peer)
	sendLastBlock(peer)
}

func makeWSURL(ip, port string) string {
	thisPort := config.GetRestPortStr()
	url := fmt.Sprintf(WSURLformat, ip, port, thisPort)

	return url
}
