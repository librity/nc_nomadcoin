package p2p

import (
	"github.com/librity/nc_nomadcoin/blockchain"
)

func BroadcastMinedBlock(minedBlock *blockchain.Block) {
	Peers.m.Lock()
	defer Peers.m.Unlock()

	for _, peer := range Peers.v {
		sendMinedBlock(peer, minedBlock)
	}
}

func BroadcastNewTx(newTx *blockchain.Tx) {
	Peers.m.Lock()
	defer Peers.m.Unlock()

	for _, peer := range Peers.v {
		sendNewTx(peer, newTx)
	}
}

func broadcastNewPeer(newPeer *peer) {
	Peers.m.Lock()
	defer Peers.m.Unlock()

	for address, receiver := range Peers.v {
		if address == newPeer.address {
			continue
		}

		sendNewPeer(receiver, newPeer.ip, newPeer.port)
	}
}
