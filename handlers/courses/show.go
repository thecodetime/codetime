package courses

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ekkapob/codetime/templates"
	"github.com/gorilla/mux"
)

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	courseName := vars["name"]
	courseFilePath := fmt.Sprintf("course/data/%v.tmpl", courseName)
	_, err := os.Stat(fmt.Sprintf("templates/%v", courseFilePath))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/courses", http.StatusTemporaryRedirect)
		return
	}

	t := templates.LoadTemplate(courseFilePath)
	t.ExecuteTemplate(w, "layout", nil)
}
