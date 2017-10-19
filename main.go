package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ekkapob/codetime/handlers"
	gh "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/namsral/flag"
)

func main() {
	var port, env string
	flag.StringVar(&port, "port", "8080", "port number")
	flag.StringVar(&env, "env", "development", "environment mode")
	flag.Parse()

	r := mux.NewRouter()
	r.StrictSlash(true)
	r.HandleFunc("/", handlers.Home)

	if env == "development" {
		r.PathPrefix("/assets/").Handler(
			http.StripPrefix("/assets", http.FileServer(http.Dir("assests"))))
	}

	port = fmt.Sprintf(":%v", port)
	fmt.Printf("server is running at port %v\n", port)
	loggedRouter := gh.LoggingHandler(os.Stdout, r)
	log.Fatal(http.ListenAndServe(port, loggedRouter))
}
