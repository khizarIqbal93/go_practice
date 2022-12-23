package main

import (
	"log"
	"net/http"
	"prep/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/crawl", handlers.RequestHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}
