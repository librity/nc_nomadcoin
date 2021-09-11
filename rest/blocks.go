package rest

import (
	"encoding/json"
	"net/http"

	"github.com/librity/nc_nomadcoin/blockchain"
)

func blocksIndex(rw http.ResponseWriter, r *http.Request) {
	blocks := blockchain.GetBlocks()

	json.NewEncoder(rw).Encode(blocks)
}

func createBlock(rw http.ResponseWriter, r *http.Request) {
	blockchain.GetBC().AddBlock()
	rw.WriteHeader(http.StatusCreated)
}

func block(rw http.ResponseWriter, r *http.Request) {
	hash := getParam(r, "hash")
	encoder := json.NewEncoder(rw)

	block, err := blockchain.FindBlock(hash)
	if err == blockchain.ErrBlockNotFound {
		rw.WriteHeader(http.StatusNotFound)
		encoder.Encode(errResp{err.Error()})
		return
	}

	encoder.Encode(block)
}
