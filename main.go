package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ekkapob/codetime/handlers"
	"github.com/namsral/flag"
)

func main() {
	var apiUrl, port, env string
	flag.StringVar(&port, "port", "4000", "port number")
	flag.StringVar(&env, "env", "development", "environment mode")
	flag.StringVar(&apiUrl, "api", "http://localhost:8080", "api url")
	flag.Parse()

	port = fmt.Sprintf(":%v", port)
	fmt.Printf("server is running at port %v\n", port)
	handler := handlers.New(env, apiUrl)
	log.Fatal(http.ListenAndServe(port, handler))
}
