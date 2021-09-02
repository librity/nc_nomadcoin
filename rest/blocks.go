package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/librity/nc_nomadcoin/blockchain"
	"github.com/librity/nc_nomadcoin/utils"
)

type newBlockBody struct {
	Data string
}

func blocksIndex(rw http.ResponseWriter, r *http.Request) {
	blocks := blockchain.GetBlockchain().GetAllBlocks()

	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(blocks)
}

func createBlock(rw http.ResponseWriter, r *http.Request) {
	var newBlock newBlockBody
	err := json.NewDecoder(r.Body).Decode(&newBlock)
	utils.HandleError(err)

	blockchain.GetBlockchain().AddBlock(newBlock.Data)
	rw.WriteHeader(http.StatusCreated)
}

func block(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	height, err := strconv.Atoi(params["height"])
	utils.HandleError(err)

	block := blockchain.GetBlockchain().GetBlock(height)
	json.NewEncoder(rw).Encode(block)
}
