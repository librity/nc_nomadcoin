package server

import (
	"net/http"

	"github.com/librity/nc_nomadcoin/blockchain"
)

type blocksData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	chain := blockchain.GetBlockchain()
	data := blocksData{"Blocks", chain.GetAllBlocks()}

	templates.ExecuteTemplate(rw, "blocks", data)
}
