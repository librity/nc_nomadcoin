package explorer

import (
	"net/http"

	"github.com/librity/nc_nomadcoin/blockchain"
)

type blocksIndexData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func blocksIndex(rw http.ResponseWriter, r *http.Request) {
	blocks := blockchain.GetBlocks()
	data := blocksIndexData{"Blocks", blocks}

	templates.ExecuteTemplate(rw, "blocks_index", data)
}
