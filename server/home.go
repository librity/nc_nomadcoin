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

var templateFunctions template.FuncMap = template.FuncMap{
	"increment": func(number int) int {
		return number + 1
	},

	"add": func(a, b int) int {
		return a + b
	},
}

func home(rw http.ResponseWriter, r *http.Request) {
	chain := blockchain.GetBlockchain()
	data := homeData{"Welcome to Nomad Coin 1.0!", chain.GetAllBlocks()}

	tmpl := template.Must(
		template.
			New("home.gohtml").
			Funcs(templateFunctions).
			ParseFiles("templates/home.gohtml"))
	tmpl.Execute(rw, data)
}
