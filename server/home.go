package server

import (
	"net/http"
	"text/template"

	"github.com/librity/nc_nomadcoin/blockchain"
)

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	chain := blockchain.GetBlockchain()
	data := homeData{"Welcome to Nomad Coin 1.0!", chain.GetAllBlocks()}

	tmpl.Execute(rw, data)
}
