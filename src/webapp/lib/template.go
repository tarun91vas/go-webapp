package lib

import "html/template"

// PopulateTemplates loads all the templates in a given directory
func PopulateTemplates(path string) *template.Template {
	templates := template.New("templates")
	template.Must(templates.ParseGlob(path + "/*.html"))
	return templates
}
