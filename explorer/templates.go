package explorer

import "text/template"

const (
	templatesDir string = "explorer/templates/"
)

var templates *template.Template

var templateFunctions template.FuncMap = template.FuncMap{
	"debug":       debug,
	"increment":   increment,
	"add":         add,
	"unixToHuman": unixToHuman,
	"homeURL":     homeURL,
	"blockURL":    blockURL,
	"txURL":       txURL,
	"walletURL":   walletURL,
}

func loadTemplates() {
	loadPages()
	loadPartials()
}

func loadPages() {
	templates = template.Must(
		template.
			New("templates").
			Funcs(templateFunctions).
			ParseGlob(templatesDir + "pages/*.gohtml"))
}

func loadPartials() {
	templates = template.Must(
		templates.
			ParseGlob(templatesDir + "partials/*.gohtml"))
}
