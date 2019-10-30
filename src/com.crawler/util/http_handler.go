package util

import (
	"strings"

	"golang.org/x/net/html"
)

const (
	HTTP = "http"
)

func Href(tk html.Token) (string, bool) {
	for _, attr := range tk.Attr {
		if attr.Key == "href" {
			return attr.Val, true
		}
	}

	return "", false
}

func IsValidDomain(url string, domains []string) bool {
	for _, domain := range domains {
		// true if starts with any valid domain
		if strings.Index(url, domain) == 0 {
			return true
		}
	}

	return false
}

func IsHttpUrl(url string) bool {
	return strings.Index(url, HTTP) == 0
}
