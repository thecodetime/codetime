package learning

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ekkapob/codetime/templates"
	"gopkg.in/yaml.v2"
)

type Author struct {
	Avatar string
	Name   string
}

type Article struct {
	Title     string
	Body      string
	Image     string
	Published string
	Link      string
	Author
}

type Content struct {
	Articles []Article
	templates.TemplateData
}

func Index(w http.ResponseWriter, r *http.Request) {
	articles, err := ioutil.ReadFile("./articles.yml")
	if err != nil {
		log.Fatal(err)
	}

	content := Content{}
	err = yaml.Unmarshal(articles, &content)
	if err != nil {
		log.Fatal(err)
	}

	t := templates.LoadTemplate("learning/index.tmpl")
	t.ExecuteTemplate(w, "layout", content)
}
