package courses

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ekkapob/codetime/templates"
)

type Course struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Name         string `json:"name"`
	Body         string `json:"body"`
	Keywords     string `json:"keywords"`
	Descriptions string `json:"descriptions"`
	Updated      int64  `json:"updated"`
}

type Courses struct {
	Data []Course `json:"data"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	apiUrl := r.Context().Value("api").(string)
	res, err := http.Get(fmt.Sprintf("%s/api/contents?type=Course", apiUrl))
	if err != nil {
		log.Fatalf("cannot get courses with error %v\n", err)
		return
	}
	defer res.Body.Close()

	var courses Courses
	json.NewDecoder(res.Body).Decode(&courses)
	t := templates.LoadTemplate("course/index.tmpl")
	t.ExecuteTemplate(w, "layout", courses)
}
