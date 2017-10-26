package handlers

import (
	"html/template"
	"path"
)

const (
	API_PATH = "http://localhost:8080"
)

func LoadTemplate(contentFile string) *template.Template {
	t := template.Must(template.ParseFiles("templates/layout/default.tmpl"))
	t = template.Must(t.ParseGlob("templates/base/*.tmpl"))
	t = template.Must(t.ParseFiles(path.Join("templates/", contentFile)))
	return t
}
