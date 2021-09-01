package explorer

import (
	"net/http"
)

type homeData struct {
	PageTitle string
}

func home(rw http.ResponseWriter, r *http.Request) {
	data := homeData{"Home"}

	templates.ExecuteTemplate(rw, "home", data)
}
