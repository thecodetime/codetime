package handlers

import (
	"context"
	"net/http"
	"os"

	"github.com/ekkapob/codetime/handlers/courses"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	apiPath = "http://localhost:8080"
)

func New(env string, apiUrl string) http.Handler {
	router := mux.NewRouter()
	router.StrictSlash(true)
	router.HandleFunc("/", Home)
	router.HandleFunc("/courses", withApi(apiUrl, courses.Index))
	if env == "development" {
		router.PathPrefix("/assets/").Handler(
			http.StripPrefix("/assets", http.FileServer(http.Dir("assets"))))
	}
	return handlers.LoggingHandler(os.Stdout, router)
}

func withApi(apiUrl string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(context.WithValue(r.Context(), "api", apiUrl))
		next(w, r)
	}
}
