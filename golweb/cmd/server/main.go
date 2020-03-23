package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	dir  = flag.String("dir", "./", "directory to serve static content")
	port = flag.String("port", "8080", "port to listen at")
)

func main() {
	flag.Parse()

	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(*dir)))
	log.Printf("Serving %s at %s", *dir, *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", *port), r); err != nil {
		log.Fatal(err)
	}
}
