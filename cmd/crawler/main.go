package main

import (
	"fmt"
	"net/http"
	"golang.org/x/net/html"
	"io"
)

func main() {
	startURLs := []string{"https://example.com"}
	for _, url := range startURLs {
		body := fetch(url)
		links := extractLinks(body)
		fmt.Println("Found links:", links)
	}
}

func fetch(url string) io.Reader {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Fetch error:", err)
		return nil
	}
	defer resp.Body.Close()
	return resp.Body

}

func extractLinks(body io.Reader) []string {
	var links []string
	tokenizer := html.NewTokenizer(body)
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			break
		}
		token := tokenizer.Token()
		if token.Data == "a" {
			for _, attr := range token.Attr {
				links = append(links, attr.Val)
			}
		}
	}
	return links
}
