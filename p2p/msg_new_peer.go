package p2p

import (
	"fmt"

	"github.com/librity/nc_nomadcoin/utils"
)

type newPeerPayload struct {
	ip, port string
}

func sendNewPeer(p *peer, ip, port string) {
	payload := newPeerPayload{
		ip:   ip,
		port: port,
	}
	message := makeMsgJSON(MsgNewPeer, payload)

	p.inbox <- message
}

func handleNewPeer(message *Msg, p *peer) {
	payload := &newPeerPayload{}
	utils.FromJSON(message.Payload, payload)

	address := buildPeerAdr(payload.ip, payload.port)
	fmt.Println("🤝 Received new peer", address, "from", p.address)
	AddPeer(payload.ip, payload.port)
}
