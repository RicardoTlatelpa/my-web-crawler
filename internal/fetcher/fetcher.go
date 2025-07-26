package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Fetcher struct {
	client *http.Client 
	userAgent string
}

func New() *Fetcher {
	return &Fetcher{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		userAgent: "GoCrawlerBot/1.0",
	}
}

func (f *Fetcher) Fetch(url string) (io.ReadCloser, error){
	req, err := http.NewRequest("Get", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create a request: %w", err)		
	}
	req.Header.Set("User-Agent", f.userAgent)

	resp, err := f.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("non-200 response: %d", resp.StatusCode)
	}
	return resp.Body, nil
}