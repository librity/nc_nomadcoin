package explorer

import "text/template"

const (
	templatesDir string = "explorer/templates/"
)

var templates *template.Template

var templateFunctions template.FuncMap = template.FuncMap{
	"increment":   increment,
	"add":         add,
	"unixToHuman": unixToHuman,
	"blockURL":    blockURL,
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
