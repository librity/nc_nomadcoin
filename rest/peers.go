package rest

import (
	"encoding/json"
	"net/http"

	"github.com/librity/nc_nomadcoin/p2p"
	"github.com/librity/nc_nomadcoin/utils"
)

type addPeerPayload struct {
	address, port string
}

func addPeer(rw http.ResponseWriter, r *http.Request) {
	payload := addPeerPayload{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	utils.PanicError(err)

	p2p.AddPeer(payload.address, payload.port)
	rw.WriteHeader(http.StatusOK)
}

func upgradeToWebSockets(rw http.ResponseWriter, r *http.Request) {
	p2p.Upgrade(rw, r)
}
