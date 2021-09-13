package p2p

import (
	"fmt"

	"github.com/gorilla/websocket"
)

var (
	Peers = make(map[string]*peer)
)

type peer struct {
	address string
	conn    *websocket.Conn
}

func (p *peer) read() {
	for {
		_, payload, err := p.conn.ReadMessage()
		if err != nil {
			break
		}

		fmt.Printf("From %s: \"%s\"\n", p.address, payload)
	}
}

func initPeer(ip, port string, conn *websocket.Conn) *peer {
	key := fmt.Sprintf("%s:%s", ip, port)
	peer := newPeer(key, conn)
	go peer.read()

	Peers[key] = peer
	return peer
}

func newPeer(address string, conn *websocket.Conn) *peer {
	peer := &peer{
		address: address,
		conn:    conn,
	}

	return peer
}
