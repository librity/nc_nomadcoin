package p2p

import "time"

func PingForever() {
	for {
		Peers.m.Lock()
		for _, peer := range Peers.v {
			peer.inbox <- []byte("Ping!")
		}
		Peers.m.Unlock()

		time.Sleep(10 * time.Second)
	}
}
