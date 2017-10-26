package courses

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ekkapob/codetime/handlers"
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
	res, err := http.Get(fmt.Sprintf("%s/api/contents?type=Course", handlers.API_PATH))
	defer res.Body.Close()
	if err != nil {
		panic("cannot retrieve courses")
		return
	}
	var courses Courses
	json.NewDecoder(res.Body).Decode(&courses)
	t := handlers.LoadTemplate("courses/index.tmpl")
	t.ExecuteTemplate(w, "layout", courses)
}
