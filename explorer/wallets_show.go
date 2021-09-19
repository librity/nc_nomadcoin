package explorer

import (
	"net/http"

	"github.com/librity/nc_nomadcoin/blockchain"
	"github.com/librity/nc_nomadcoin/utils"
)

type walletsShowData struct {
	PageTitle     string
	Address       string
	Balance       uint
	UnspTxOutputs []*blockchain.UnspTxOutput
}

func walletsShow(rw http.ResponseWriter, r *http.Request) {
	address := utils.GetRoute(r, "address")
	outputs := blockchain.UnspTxOutputsFrom(address)
	balance := blockchain.SumOverBalance(outputs)

	data := walletsShowData{
		PageTitle:     "Show Wallet",
		Address:       address,
		Balance:       balance,
		UnspTxOutputs: outputs,
	}
	templates.ExecuteTemplate(rw, "wallets_show", data)
}
