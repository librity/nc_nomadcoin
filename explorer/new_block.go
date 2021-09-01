package explorer

import (
	"net/http"
)

type newBlockData struct {
	PageTitle string
}

func newBlock(rw http.ResponseWriter, r *http.Request) {
	data := newBlockData{"New Block"}

	templates.ExecuteTemplate(rw, "new_block", data)
}
