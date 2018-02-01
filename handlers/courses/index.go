package courses

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ekkapob/codetime/templates"
	yaml "gopkg.in/yaml.v2"
)

// type Course struct {
// 	Id           int    `json:"id"`
// 	Title        string `json:"title"`
// 	Name         string `json:"name"`
// 	Body         string `json:"body"`
// 	Keywords     string `json:"keywords"`
// 	Descriptions string `json:"descriptions"`
// 	Updated      int64  `json:"updated"`
// }

// type Courses struct {
// 	Data []Course `json:"data"`
// 	templates.TemplateData
// }

type Author struct {
	Avatar string
	Name   string
	Skill  string
}

type Course struct {
	Title string
	Body  string
	Image string
	Link  string
	Alt   string
	Author
}

type Content struct {
	Courses []Course
	templates.TemplateData
}

func Index(w http.ResponseWriter, r *http.Request) {
	// apiUrl := r.Context().Value("api").(string)
	// res, err := http.Get(fmt.Sprintf("%s/api/contents?type=Course", apiUrl))
	// if err != nil {
	// 	log.Fatalf("cannot get courses with error %v\n", err)
	// 	return
	// }
	// defer res.Body.Close()
	// var courses Courses
	// json.NewDecoder(res.Body).Decode(&courses)

	articles, err := ioutil.ReadFile("./courses.yml")
	if err != nil {
		log.Fatal(err)
	}

	content := Content{}
	err = yaml.Unmarshal(articles, &content)
	if err != nil {
		log.Fatal(err)
	}
	t := templates.LoadTemplate("course/index.tmpl")
	t.ExecuteTemplate(w, "layout", content)
}
