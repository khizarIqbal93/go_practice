package crawl

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

type Page struct {
	PageUrl   *url.URL `json:"-"`
	Page      string   `json:"page"`
	ParentUrl *url.URL `json:"-"`
	Parent    string   `json:"parent"`
	Links     []Page   `json:"links,omitempty"`
}

func (p *Page) SetPageUrl(urlString string, isRoot bool) {
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		log.Println("url could not be parsed")
	}

	if parsedUrl.Scheme == "" {
		parsedUrl.Scheme = "https"
	}

	// TODO fix this
	if parsedUrl.Host == "" {
		newHost, newPath, _ := strings.Cut(parsedUrl.Path, "/")
		parsedUrl.Host = newHost
		if newHost == "" {
			parsedUrl.Host = p.ParentUrl.Host
		}
		parsedUrl.Path = "/" + newPath
	}

	p.PageUrl = parsedUrl
	if isRoot {
		p.ParentUrl, err = url.Parse(p.PageUrl.String())
		if err != nil {
			log.Println("something went wrong when parsing")
		}
	}
	p.Page = p.PageUrl.String()
	p.Parent = p.ParentUrl.String()
}

// returns the DOM of urlString as a string
func getHtml(urlString string) string {
	resp, err := http.Get(urlString)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	html := string(body)
	return html
}

func (p *Page) GetLinks(visited map[string]int) {
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					child := Page{}
					child.ParentUrl = p.PageUrl
					child.SetPageUrl(a.Val, false)
					p.Links = append(p.Links, child)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	if visited[p.PageUrl.String()] == 0 && p.ParentUrl.Host == p.PageUrl.Host {
		doc, err := html.Parse(strings.NewReader(getHtml(p.PageUrl.String())))
		if err != nil {
			log.Println("Error parsing HTML")
		}
		visited[p.PageUrl.String()]++
		f(doc)
	}
}
