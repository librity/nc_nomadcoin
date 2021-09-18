package explorer

import (
	"net/http"

	"github.com/librity/nc_nomadcoin/blockchain"
	"github.com/librity/nc_nomadcoin/utils"
)

type txsShowData struct {
	PageTitle string
	Tx        *blockchain.Tx
}

func txsShow(rw http.ResponseWriter, r *http.Request) {
	id := utils.GetRoute(r, "id")
	tx, err := blockchain.FindTx(id)
	if err == blockchain.ErrTxNotFound {
		http.Redirect(rw, r, "/404", http.StatusFound)
	}

	data := txsShowData{"Transaction", tx}
	templates.ExecuteTemplate(rw, "txs_show", data)
}
