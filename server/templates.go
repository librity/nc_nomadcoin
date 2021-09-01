package server

import "text/template"

const (
	templatesDir string = "templates/"
)

var templateFunctions template.FuncMap = template.FuncMap{
	"increment": func(number int) int {
		return number + 1
	},

	"add": func(a, b int) int {
		return a + b
	},
}

func loadTemplates() {
	templates = template.Must(template.New("dsada").Funcs(templateFunctions).ParseGlob(templatesDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templatesDir + "partials/*.gohtml"))
	// templates = templates.
}
