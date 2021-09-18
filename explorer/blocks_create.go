package explorer

import (
	"net/http"

	"github.com/librity/nc_nomadcoin/blockchain"
)

func blocksCreate(rw http.ResponseWriter, r *http.Request) {
	blockchain.MineBlock()

	http.Redirect(rw, r, "/blocks", http.StatusFound)
}
