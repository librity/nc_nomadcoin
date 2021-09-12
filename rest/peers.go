package rest

import (
	"encoding/json"
	"net/http"

	"github.com/librity/nc_nomadcoin/p2p"
	"github.com/librity/nc_nomadcoin/utils"
)

type addPeerPayload struct {
	IP, Port string
}

func peersIndex(rw http.ResponseWriter, r *http.Request) {
	peers := p2p.Peers
	json.NewEncoder(rw).Encode(peers)
}

func addPeer(rw http.ResponseWriter, r *http.Request) {
	payload := addPeerPayload{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	utils.PanicError(err)

	thisPort := cleanPort()
	p2p.AddPeer(payload.IP, payload.Port, thisPort)
	rw.WriteHeader(http.StatusOK)
}

func upgradePeer(rw http.ResponseWriter, r *http.Request) {
	p2p.UpgradePeer(rw, r)
}
