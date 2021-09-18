package explorer

import "net/http"

const (
	staticDir   string = "explorer/static"
	staticRoute string = "/static/"
)

var (
	router *http.ServeMux = http.NewServeMux()
)

func loadFileServer() {
	fileServer := http.FileServer(http.Dir(staticDir))
	router.Handle(staticRoute, http.StripPrefix(staticRoute, fileServer))
}

func loadRoutes() {
	router.HandleFunc("/", home)

	router.HandleFunc("/blocks", blocks)
	router.HandleFunc("/blocks/mine", mineBlock)
}
