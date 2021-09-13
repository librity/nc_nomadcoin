package p2p

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type peer struct {
	ip      string
	port    string
	address string
	conn    *websocket.Conn
	inbox   chan []byte
}

func (p *peer) read() {
	defer p.remove()

	for {
		message := &Msg{}
		err := p.conn.ReadJSON(message)
		if err != nil {
			break
		}

		fmt.Println("Received from", p.address, message)
	}
}

func (p *peer) write() {
	defer p.remove()

	for {
		payload, ok := <-p.inbox
		if !ok {
			break
		}

		err := p.conn.WriteMessage(websocket.TextMessage, payload)
		if err != nil {
			break
		}

		fmt.Printf("Sent to %s: \"%s\"\n---\n", p.address, payload)
	}
}

func (p *peer) remove() {
	p.conn.Close()
	delistPeer(p)
}

func initPeer(ip, port string, conn *websocket.Conn) *peer {
	peer := newPeer(ip, port, conn)
	go peer.read()
	go peer.write()

	insertPeer(peer)
	return peer
}

func newPeer(ip, port string, conn *websocket.Conn) *peer {
	address := fmt.Sprintf("%s:%s", ip, port)
	peer := &peer{
		ip:      ip,
		port:    port,
		address: address,
		conn:    conn,
		inbox:   make(chan []byte),
	}

	return peer
}
