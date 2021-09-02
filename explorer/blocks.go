package explorer

import (
	"net/http"

	"github.com/librity/nc_nomadcoin/blockchain"
)

type blocksData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		blocksIndex(rw, r)
	case "POST":
		createBlock(rw, r)
	}
}

func blocksIndex(rw http.ResponseWriter, r *http.Request) {
	blocks := blockchain.Get().AllBlocks()
	data := blocksData{"Blocks", blocks}

	templates.ExecuteTemplate(rw, "blocks", data)
}

func createBlock(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	blockData := r.Form.Get("blockData")
	blockchain.Get().AddBlock(blockData)

	http.Redirect(rw, r, "/blocks", http.StatusFound)
}
