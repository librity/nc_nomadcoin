package rest

import (
	"encoding/json"
	"net/http"

	"github.com/librity/nc_nomadcoin/blockchain"
	"github.com/librity/nc_nomadcoin/utils"
)

type wltResp struct {
	Address string `json:"address"`
	Balance uint   `json:"balance"`
}

type wltDetailsResp struct {
	Address        string                        `json:"address"`
	Balance        uint                          `json:"balance"`
	UnspentOutputs []*blockchain.UnspentTxOutput `json:"unspentOutputs"`
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
	balance := blockchain.BalanceOf(address)

	response := wltResp{
		Address: address,
		Balance: balance,
	}
	err := json.NewEncoder(rw).Encode(response)
	utils.HandleError(err)
}

func handleFullInfo(rw http.ResponseWriter, r *http.Request) {
	address := getParam(r, "address")
	outputs := blockchain.UnspentTxOutputsFrom(address)
	balance := blockchain.SumOverBalance(outputs)

	response := wltDetailsResp{
		Address:        address,
		Balance:        balance,
		UnspentOutputs: outputs,
	}
	err := json.NewEncoder(rw).Encode(response)
	utils.HandleError(err)
}
