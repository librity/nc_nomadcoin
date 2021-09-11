package rest

import (
	"net/http"

	"github.com/librity/nc_nomadcoin/p2p"
)

func upgradeToWebSockets(rw http.ResponseWriter, r *http.Request) {
	p2p.Upgrade(rw, r)
}
