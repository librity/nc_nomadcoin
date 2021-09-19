package explorer

import (
	"net/http"

	"github.com/librity/nc_nomadcoin/blockchain"
	"github.com/librity/nc_nomadcoin/wallet"
)

type walletsServerData struct {
	PageTitle     string
	Address       string
	Balance       uint
	UnspTxOutputs []*blockchain.UnspTxOutput
}

func walletsServer(rw http.ResponseWriter, r *http.Request) {
	address := wallet.GetAddress()
	outputs := blockchain.UnspTxOutputsFrom(address)
	balance := blockchain.SumOverBalance(outputs)

	data := walletsServerData{
		PageTitle:     "Show Wallet",
		Address:       address,
		Balance:       balance,
		UnspTxOutputs: outputs,
	}
	templates.ExecuteTemplate(rw, "wallets_show", data)
}
