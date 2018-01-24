package handlers

import (
	"net/http"

	"github.com/ekkapob/codetime/templates"
)

func OurWork(w http.ResponseWriter, r *http.Request) {
	t := templates.LoadTemplate("our_work.tmpl")
	t.ExecuteTemplate(w, "layout", templates.TemplateData{})
}
