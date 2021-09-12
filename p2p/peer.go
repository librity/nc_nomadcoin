package p2p

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type peer struct {
	conn *websocket.Conn
}

var (
	Peers = make(map[string]*peer)
)

func initPeer(ip, port string, conn *websocket.Conn) *peer {
	peer := newPeer(conn)
	key := fmt.Sprintf("%s:%s", ip, port)
	Peers[key] = peer

	return peer
}

func newPeer(conn *websocket.Conn) *peer {
	peer := &peer{
		conn: conn,
	}

	return peer
}