package handlers

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(baseTemplate().ParseFiles("templates/home.tmpl"))
	t.ExecuteTemplate(w, "layout", nil)
}
