package handlers

import (
	"net/http"

	"github.com/ekkapob/codetime/templates"
)

func Contact(w http.ResponseWriter, r *http.Request) {
	t := templates.LoadTemplate("contact.tmpl")
	t.ExecuteTemplate(w, "layout", templates.TemplateData{})
}
