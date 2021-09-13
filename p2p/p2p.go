package p2p

import "time"

func PingForever() {
	for {
		Peers.mx.Lock()
		for _, peer := range Peers.active {
			peer.inbox <- []byte("Ping!")
		}
		Peers.mx.Unlock()

		time.Sleep(10 * time.Second)
	}
}
