package rest

import (
	"encoding/json"
	"net/http"

	"github.com/librity/nc_nomadcoin/blockchain"
	"github.com/librity/nc_nomadcoin/p2p"
	"github.com/librity/nc_nomadcoin/utils"
)

func blocksIndex(rw http.ResponseWriter, r *http.Request) {
	blocks := blockchain.GetBlocks()

	json.NewEncoder(rw).Encode(blocks)
}

func createBlock(rw http.ResponseWriter, r *http.Request) {
	newBlock := blockchain.GetBC().AddBlock()
	p2p.BroadcastMinedBlock(newBlock)

	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(newBlock)
}

func block(rw http.ResponseWriter, r *http.Request) {
	hash := utils.GetParam(r, "hash")
	encoder := json.NewEncoder(rw)

	block, err := blockchain.FindBlock(hash)
	if err == blockchain.ErrBlockNotFound {
		rw.WriteHeader(http.StatusNotFound)
		encoder.Encode(errResp{err.Error()})
		return
	}

	encoder.Encode(block)
}
