package templates

import (
	"html/template"
	"path"
)

type TemplateData struct {
	NoHeader bool
}

func LoadTemplate(contentFile string) *template.Template {
	t := template.Must(template.ParseFiles("templates/layout/default.tmpl"))
	t = template.Must(t.ParseGlob("templates/base/*.tmpl"))
	t = template.Must(t.ParseFiles(path.Join("templates/", contentFile)))
	return t
}
