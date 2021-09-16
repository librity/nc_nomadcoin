package p2p

import (
	"fmt"

	"github.com/librity/nc_nomadcoin/utils"
)

type newPeerPayload struct {
	IP, Port string
}

func sendNewPeer(p *peer, ip, port string) {
	payload := newPeerPayload{
		IP:   ip,
		Port: port,
	}
	message := makeMsgJSON(MsgNewPeer, payload)

	p.inbox <- message
}

func handleNewPeer(message *Msg, p *peer) {
	payload := &newPeerPayload{}
	utils.FromJSON(message.Payload, payload)

	address := buildPeerAdr(payload.IP, payload.Port)
	if badPeerPayload(payload) {
		fmt.Println("🤝 Received invalid peer", address, "from", p.address)
		return
	}

	fmt.Println("🤝 Received new peer", address, "from", p.address)
	AddPeer(payload.IP, payload.Port)
}

func badPeerPayload(payload *newPeerPayload) bool {
	isBad := payload.IP == "" || payload.Port == ""
	// TODO: Better peer validation

	return isBad
}
