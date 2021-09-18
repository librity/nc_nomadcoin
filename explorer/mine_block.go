package explorer

import (
	"net/http"
)

type mineBlockData struct {
	PageTitle string
}

func mineBlock(rw http.ResponseWriter, r *http.Request) {
	data := mineBlockData{"Mine Block"}

	templates.ExecuteTemplate(rw, "mine_block", data)
}
