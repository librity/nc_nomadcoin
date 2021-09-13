package p2p

import "time"

func PingForever() {
	for {
		for _, peer := range Peers {
			peer.inbox <- []byte("Ping!")
		}

		time.Sleep(10 * time.Second)
	}
}
