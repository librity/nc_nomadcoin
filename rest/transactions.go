package rest

import (
	"encoding/json"
	"net/http"

	"github.com/librity/nc_nomadcoin/blockchain"
	"github.com/librity/nc_nomadcoin/p2p"
	"github.com/librity/nc_nomadcoin/utils"
)

type txPayload struct {
	To     string
	Amount uint
}

func createTx(rw http.ResponseWriter, r *http.Request) {
	payload := txPayload{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	utils.PanicError(err)

	tx, err := blockchain.AddTx(payload.To, payload.Amount)
	if err == blockchain.ErrNotEnoughMoney {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(errResp{err.Error()})
		return
	}

	go p2p.BroadcastNewTx(tx)
	rw.WriteHeader(http.StatusCreated)
}
