package explorer

import (
	"net/http"

	"github.com/librity/nc_nomadcoin/blockchain"
)

type blocksData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func blocksIndex(rw http.ResponseWriter, r *http.Request) {
	blocks := blockchain.GetBlocks()
	data := blocksData{"Blocks", blocks}

	templates.ExecuteTemplate(rw, "blocks", data)
}
