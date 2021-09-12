package rest

import (
	"encoding/json"
	"net/http"

	"github.com/librity/nc_nomadcoin/blockchain"
	"github.com/librity/nc_nomadcoin/utils"
	"github.com/librity/nc_nomadcoin/wallet"
)

type wltResp struct {
	Address string `json:"address"`
	Balance uint   `json:"balance"`
}

type wltDetailsResp struct {
	Address      string                     `json:"address"`
	Balance      uint                       `json:"balance"`
	UnspTxOutput []*blockchain.UnspTxOutput `json:"unspTxOutputs"`
}

func serverWlt(rw http.ResponseWriter, r *http.Request) {
	details := utils.GetQuery(r, "details")
	address := wallet.GetAddress()

	switch details {
	case "true":
		handleFullInfo(rw, r, address)
	default:
		handleBalanceOnly(rw, r, address)
	}
}

func wlt(rw http.ResponseWriter, r *http.Request) {
	details := utils.GetQuery(r, "details")
	address := utils.GetParam(r, "address")

	switch details {
	case "true":
		handleFullInfo(rw, r, address)
	default:
		handleBalanceOnly(rw, r, address)
	}
}

func handleBalanceOnly(rw http.ResponseWriter, r *http.Request, address string) {
	balance := blockchain.BalanceOf(address)

	response := wltResp{
		Address: address,
		Balance: balance,
	}
	err := json.NewEncoder(rw).Encode(response)
	utils.PanicError(err)
}

func handleFullInfo(rw http.ResponseWriter, r *http.Request, address string) {
	outputs := blockchain.UnspTxOutputsFrom(address)
	balance := blockchain.SumOverBalance(outputs)

	response := wltDetailsResp{
		Address:      address,
		Balance:      balance,
		UnspTxOutput: outputs,
	}
	err := json.NewEncoder(rw).Encode(response)
	utils.PanicError(err)
}
