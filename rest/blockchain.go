package rest

import (
	"encoding/json"
	"net/http"

	"github.com/librity/nc_nomadcoin/blockchain"
)

func blokchainStatus(rw http.ResponseWriter, r *http.Request) {
	bc := blockchain.Get()
	json.NewEncoder(rw).Encode(bc)
}
