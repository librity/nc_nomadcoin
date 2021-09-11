package rest

import (
	"encoding/json"
	"net/http"

	"github.com/librity/nc_nomadcoin/blockchain"
	"github.com/librity/nc_nomadcoin/utils"
)

type txPayload struct {
	To     string
	Amount uint
}

func createTransaction(rw http.ResponseWriter, r *http.Request) {
	payload := txPayload{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	utils.HandleError(err)

	err = blockchain.Mempool.AddTx(payload.To, payload.Amount)
	if err == blockchain.ErrNotEnoughMoney {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(errResp{err.Error()})
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
