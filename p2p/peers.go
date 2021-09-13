package p2p

import (
	"fmt"
	"sync"
)

var (
	Peers = &peers{
		active: make(map[string]*peer),
	}
)

type peers struct {
	active map[string]*peer
	mx     sync.Mutex
}

func GetPeers() []string {
	Peers.mx.Lock()
	defer Peers.mx.Unlock()

	var peersList []string
	for address := range Peers.active {
		peersList = append(peersList, address)
	}
	return peersList
}

func insertPeer(p *peer) {
	Peers.mx.Lock()
	fmt.Println("Inserting peer", p.address)
	defer Peers.mx.Unlock()

	Peers.active[p.address] = p
}

func delistPeer(p *peer) {
	Peers.mx.Lock()
	fmt.Println("Delisting peer", p.address)
	defer Peers.mx.Unlock()

	delete(Peers.active, p.address)
}
