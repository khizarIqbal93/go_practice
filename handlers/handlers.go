package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"prep/pkg/crawl"
)

type CrawlRequest struct {
	Url   string `json:"url"`
	Depth int    `json:"depth"`
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Something is wrong with the body of the request")
		log.Println(err)
	}

	data := CrawlRequest{}
	json.Unmarshal(body, &data)

	visited := make(map[string]int)

	// initialize parent page
	page := crawl.Page{}
	page.SetPageUrl(data.Url, true)

	page.GetLinks(visited)
	res, err := json.Marshal(page)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "something went wrong")
	}

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, string(res))
}
