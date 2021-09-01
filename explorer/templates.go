package explorer

import "text/template"

const (
	templatesDir string = "explorer/templates/"
)

var templates *template.Template

var templateFunctions template.FuncMap = template.FuncMap{
	"increment": func(number int) int {
		return number + 1
	},

	"add": func(a, b int) int {
		return a + b
	},
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
