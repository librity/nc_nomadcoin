package rest

import (
	"encoding/json"
	"net/http"

	"github.com/librity/nc_nomadcoin/blockchain"
)

func blokchain(rw http.ResponseWriter, r *http.Request) {
	bc := blockchain.GetBC()
	json.NewEncoder(rw).Encode(bc)
}
