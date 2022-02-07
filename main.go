package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type MTripper struct {
	rt http.RoundTripper
}

func (m MTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	fmt.Println("start")
	now := time.Now()
	defer func() {
		elapsed := time.Since(now)
		fmt.Println("request took, ", elapsed)
	}()
	return m.rt.RoundTrip(request)

}

func main() {
	mt := MTripper{
		http.DefaultTransport,
	}
	c := http.Client{
		Transport: mt,
	}
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://example.com", nil)
	if err != nil {
		log.Fatalf("error: %e", err)
	}
	_, _ = c.Do(req)
}
