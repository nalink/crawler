package service

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"

	"com.crawler/util"
)

const (
	ANCHOR = "a"
)

type Crawler struct {
	domains []string
}

func NewCrawler(domains []string) *Crawler {
	return &Crawler{domains: domains}
}

func (cl *Crawler) Crawl(url string, links chan string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("\nUnable to crawl '%s': %s", url, err)
		return
	}

	raw := resp.Body
	if raw == nil {
		fmt.Printf("\nCan not read the page '%s'", url)
		return
	}

	defer raw.Close()

	tk := html.NewTokenizer(raw)

	for {
		tn := tk.Next()

		switch {
		case tn == html.ErrorToken:
			return
		case tn == html.StartTagToken:
			token := tk.Token()

			if token.Data != ANCHOR {
				continue
			}

			href, ok := util.Href(token)
			if !ok {
				continue
			}

			if util.IsValidDomain(href, cl.domains) {
				links <- href
				go cl.Crawl(href, links) // starting new go routines
			} else if util.IsHttpUrl(href) { // ignoring non-http urls
				fmt.Printf("\n[Non-Domain] %s", href)
			}
		}
	}
}
