package p2p

import (
	"fmt"
	"sync"
)

var (
	Peers = &peers{
		v: make(map[string]*peer),
	}
)

// Abreviations per convention
type peers struct {
	v map[string]*peer
	m sync.Mutex
}

func GetPeersList() []string {
	Peers.m.Lock()
	defer Peers.m.Unlock()

	var peersList []string
	for address := range Peers.v {
		peersList = append(peersList, address)
	}
	return peersList
}

func insertPeer(p *peer) {
	Peers.m.Lock()
	defer Peers.m.Unlock()

	fmt.Println("🤝 Inserting peer", p.address)
	Peers.v[p.address] = p
}

func delistPeer(p *peer) {
	Peers.m.Lock()
	defer Peers.m.Unlock()

	fmt.Println("🤝 Delisting peer", p.address)
	delete(Peers.v, p.address)
}
