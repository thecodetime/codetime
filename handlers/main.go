package handlers

import "html/template"

func baseTemplate() *template.Template {
	t, err := template.ParseFiles("templates/layout/default.tmpl")
	if err != nil {
		panic("cannot render template")
	}
	t, _ = t.ParseGlob("templates/base/*.tmpl")
	if err != nil {
		panic("cannot render template")
	}
	return t
}
