package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/librity/nc_nomadcoin/blockchain"
	"github.com/librity/nc_nomadcoin/utils"
)

type newBlockBody struct {
	Data string
}

// func blocksIndex(rw http.ResponseWriter, r *http.Request) {
// 	blocks := blockchain.Get().AllBlocks()

// 	json.NewEncoder(rw).Encode(blocks)
// }

func createBlock(rw http.ResponseWriter, r *http.Request) {
	var newBlock newBlockBody
	err := json.NewDecoder(r.Body).Decode(&newBlock)
	utils.HandleError(err)

	blockchain.Get().AddBlock(newBlock.Data)
	rw.WriteHeader(http.StatusCreated)
}

func block(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	hash := params["hash"]

	encoder := json.NewEncoder(rw)
	block, err := blockchain.FindBlock(hash)
	if err == blockchain.ErrBlockNotFound {
		rw.WriteHeader(http.StatusNotFound)
		message := fmt.Sprint(err)
		encoder.Encode(errorResponse{message})
		return
	}

	encoder.Encode(block)
}
