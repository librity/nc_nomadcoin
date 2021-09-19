package explorer

import (
	"net/http"

	"github.com/librity/nc_nomadcoin/blockchain"
	"github.com/librity/nc_nomadcoin/utils"
)

type blocksShowData struct {
	PageTitle string
	Block     *blockchain.Block
}

func blocksShow(rw http.ResponseWriter, r *http.Request) {
	hash := utils.GetRoute(r, "hash")
	block, err := blockchain.FindBlock(hash)
	if err == blockchain.ErrBlockNotFound {
		http.Redirect(rw, r, "/404", http.StatusFound)
	}

	data := blocksShowData{"Show Block", block}
	templates.ExecuteTemplate(rw, "blocks_show", data)
}
