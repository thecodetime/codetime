package handlers

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/layout.html"))
	t.Execute(w, nil)
}
