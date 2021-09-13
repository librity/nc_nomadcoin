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
	inbox   chan []byte
}

func (p *peer) read() {
	for {
		_, payload, err := p.conn.ReadMessage()
		if err != nil {
			break
		}

		fmt.Printf("Received from %s:\n\"%s\"\n---\n", p.address, payload)
	}
}

func (p *peer) write() {
	for {
		payload, ok := <-p.inbox
		if !ok {
			break
		}
		err := p.conn.WriteMessage(websocket.TextMessage, payload)
		if err != nil {
			break
		}

		fmt.Printf("Sent to %s:\n\"%s\"\n---\n", p.address, payload)
	}
}

func initPeer(ip, port string, conn *websocket.Conn) *peer {
	key := fmt.Sprintf("%s:%s", ip, port)
	peer := newPeer(key, conn)
	go peer.read()
	go peer.write()

	Peers[key] = peer
	return peer
}

func newPeer(address string, conn *websocket.Conn) *peer {
	peer := &peer{
		address: address,
		conn:    conn,
		inbox:   make(chan []byte),
	}

	return peer
}
