package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type CrawlRequest struct {
	Url   string `json:"name"`
	Depth int    `json:"depth"`
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	data := CrawlRequest{}
	json.Unmarshal(body, &data)
	w.WriteHeader(http.StatusCreated)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/category", requestHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}
