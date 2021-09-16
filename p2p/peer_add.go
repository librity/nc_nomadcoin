package p2p

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/librity/nc_nomadcoin/config"
	"github.com/librity/nc_nomadcoin/utils"
)

const (
	WSURLformat = "ws://%s:%s/peers/upgrade?senderPort=%s"
)

func AddPeer(ip, port string) {
	address := buildPeerAdr(ip, port)
	fmt.Println("🤝 Adding peer", address)

	Peers.m.Lock()
	_, peerExists := Peers.v[address]
	Peers.m.Unlock()
	if peerExists {
		fmt.Println("🤝 Already connected to peer", address)
		return
	}

	addPeer(ip, port)
}

func addPeer(ip, port string) {
	url := makeWSURL(ip, port)
	seniorConn, _, err := websocket.DefaultDialer.Dial(url, nil)
	utils.PanicError(err)

	peer := initPeer(ip, port, seniorConn)
	go broadcastNewPeer(peer)
	sendLastBlock(peer)
}

func makeWSURL(ip, port string) string {
	senderPort := config.GetRestPortStr()
	url := fmt.Sprintf(WSURLformat, ip, port, senderPort)

	return url
}
