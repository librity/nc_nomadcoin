package rest

import (
	"encoding/json"
	"net/http"

	"github.com/librity/nc_nomadcoin/blockchain"
	"github.com/librity/nc_nomadcoin/utils"
)

type balanceResponse struct {
	Balance   int                    `json:"balance"`
	TxOutputs []*blockchain.TxOutput `json:"transactionOuptuts"`
}

func wallet(rw http.ResponseWriter, r *http.Request) {
	address := getParam(r, "address")

	chain := blockchain.Get()
	outputs := chain.TxOutputsFrom(address)
	balance := blockchain.SumOverBalance(outputs)

	response := balanceResponse{
		Balance:   balance,
		TxOutputs: outputs,
	}
	err := json.NewEncoder(rw).Encode(response)
	utils.HandleError(err)
}
