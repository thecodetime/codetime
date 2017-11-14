package handlers

import (
	"net/http"

	"github.com/ekkapob/codetime/templates"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t := templates.LoadTemplate("home.tmpl")
	t.ExecuteTemplate(w, "layout", templates.TemplateData{true})
}
