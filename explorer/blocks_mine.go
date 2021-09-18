package explorer

import (
	"net/http"
)

type blocksMineData struct {
	PageTitle string
}

func blocksMine(rw http.ResponseWriter, r *http.Request) {
	data := blocksMineData{"Mine Block"}

	templates.ExecuteTemplate(rw, "blocks_mine", data)
}
