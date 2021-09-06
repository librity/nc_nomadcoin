package rest

import (
	"encoding/json"
	"net/http"

	"github.com/librity/nc_nomadcoin/blockchain"
	"github.com/librity/nc_nomadcoin/utils"
)

func mempool(rw http.ResponseWriter, r *http.Request) {
	mempool := blockchain.Mempool

	err := json.NewEncoder(rw).Encode(mempool)
	utils.HandleError(err)
}
