package server

import (
	"net/http"

	"github.com/librity/nc_nomadcoin/blockchain"
)

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	chain := blockchain.GetBlockchain()
	data := homeData{"Welcome to Nomad Coin 1.0!", chain.GetAllBlocks()}

	templates.ExecuteTemplate(rw, "home", data)
}
