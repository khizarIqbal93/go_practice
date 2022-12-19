package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type CrawlRequest struct {
	Url   string `json:"url"`
	Depth int    `json:"depth"`
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, http.StatusBadRequest)
		log.Println(err)
	}

	data := CrawlRequest{}
	json.Unmarshal(body, &data)
	log.Println(data)
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "crawling")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/crawl", requestHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}
