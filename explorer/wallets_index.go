package explorer

import (
	"net/http"

	"github.com/librity/nc_nomadcoin/blockchain"
)

type walletsIndexData struct {
	PageTitle string
	Addresses []string
}

func walletsIndex(rw http.ResponseWriter, r *http.Request) {
	addresses := blockchain.GetAddresses()
	data := walletsIndexData{"Wallets", addresses}

	templates.ExecuteTemplate(rw, "wallets_index", data)
}
