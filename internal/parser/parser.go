package parser

import (
	"io"
	"net/url"

	"golang.org/x/net/html"
)

func ExtractLinks(baseURL string, body io.Reader)([]string, error){
	var links []string

	tokenizer := html.NewTokenizer(body)

	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			break // End of document
		}
		token := tokenizer.Token()
		// only interested in <a> tags
		if token.Data == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					href := attr.Val
					absoluteURL := resolveURL(baseURL, href)
					if absoluteURL != "" {
						links = append(links, absoluteURL)
					}
				}
			}
		}
	}
	return links, nil
}

func resolveURL(base string, href string) string {
	baseURL, err := url.Parse(base)
	if err != nil {
		return ""
	}
	hrefURL, err := url.Parse(href)
	if err != nil {
		return ""
	}

	resolved := baseURL.ResolveReference(hrefURL)
	return resolved.String()
}