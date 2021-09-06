package rest

import (
	"encoding/json"
	"net/http"

	"github.com/librity/nc_nomadcoin/blockchain"
	"github.com/librity/nc_nomadcoin/utils"
)

type wltResp struct {
	Address string `json:"address"`
	Balance int    `json:"balance"`
}

type wltDetailsResp struct {
	Address   string                 `json:"address"`
	Balance   int                    `json:"balance"`
	TxOutputs []*blockchain.TxOutput `json:"transactionOuptuts"`
}

func wallet(rw http.ResponseWriter, r *http.Request) {
	details := getQuery(r, "details")

	switch details {
	case "true":
		handleFullInfo(rw, r)
	default:
		handleBalanceOnly(rw, r)
	}
}

func handleBalanceOnly(rw http.ResponseWriter, r *http.Request) {
	address := getParam(r, "address")
	chain := blockchain.Get()
	balance := chain.BalanceOf(address)

	response := wltResp{
		Address: address,
		Balance: balance,
	}
	err := json.NewEncoder(rw).Encode(response)
	utils.HandleError(err)
}

func handleFullInfo(rw http.ResponseWriter, r *http.Request) {
	address := getParam(r, "address")
	chain := blockchain.Get()
	outputs := chain.TxOutputsFrom(address)
	balance := blockchain.SumOverBalance(outputs)

	response := wltDetailsResp{
		Address:   address,
		Balance:   balance,
		TxOutputs: outputs,
	}
	err := json.NewEncoder(rw).Encode(response)
	utils.HandleError(err)
}
