package rest

import (
	"encoding/json"
	"net/http"

	"github.com/librity/nc_nomadcoin/blockchain"
	"github.com/librity/nc_nomadcoin/utils"
)

type NewBlockBody struct {
	Data string
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		blocksIndex(rw, r)
	case "POST":
		createBlock(rw, r)
	}
}

func blocksIndex(rw http.ResponseWriter, r *http.Request) {
	blocks := blockchain.GetBlockchain().GetAllBlocks()

	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(blocks)
}

func createBlock(rw http.ResponseWriter, r *http.Request) {
	var newBlockBody NewBlockBody
	err := json.NewDecoder(r.Body).Decode(&newBlockBody)
	utils.HandleError(err)

	blockchain.GetBlockchain().AddBlock(newBlockBody.Data)
	rw.WriteHeader(http.StatusCreated)
}
