package handlers

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t := LoadTemplate("home.tmpl")
	t.ExecuteTemplate(w, "layout", nil)
}
