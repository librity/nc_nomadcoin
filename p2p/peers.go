package p2p

import (
	"fmt"
	"sync"

	"github.com/librity/nc_nomadcoin/blockchain"
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

func BroadcastMinedBlock(minedBlock *blockchain.Block) {
	Peers.m.Lock()
	defer Peers.m.Unlock()

	for _, peer := range Peers.v {
		sendMinedBlock(peer, minedBlock)
	}
}

func insertPeer(p *peer) {
	Peers.m.Lock()
	fmt.Println("🤝 Inserting peer", p.address)
	defer Peers.m.Unlock()

	Peers.v[p.address] = p
}

func delistPeer(p *peer) {
	Peers.m.Lock()
	fmt.Println("🤝 Delisting peer", p.address)
	defer Peers.m.Unlock()

	delete(Peers.v, p.address)
}
