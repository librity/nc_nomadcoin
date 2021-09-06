package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/librity/nc_nomadcoin/blockchain"
)

func blocksIndex(rw http.ResponseWriter, r *http.Request) {
	blocks := blockchain.Get().Blocks()

	json.NewEncoder(rw).Encode(blocks)
}

func createBlock(rw http.ResponseWriter, r *http.Request) {
	blockchain.Get().AddBlock()
	rw.WriteHeader(http.StatusCreated)
}

func block(rw http.ResponseWriter, r *http.Request) {
	hash := getParam(r, "hash")
	encoder := json.NewEncoder(rw)

	block, err := blockchain.FindBlock(hash)
	if err == blockchain.ErrBlockNotFound {
		rw.WriteHeader(http.StatusNotFound)
		message := fmt.Sprint(err)
		encoder.Encode(errResp{message})
		return
	}

	encoder.Encode(block)
}
