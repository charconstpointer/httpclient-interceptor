package fleet

import (
	"fmt"
	"net/http"
	"time"
)

type Middleware func(*http.Request)

type MTripper struct {
	rt          http.RoundTripper
	middlewares []Middleware
}

func (m *MTripper) Use(middleware Middleware) {
	m.middlewares = append(m.middlewares, middleware)
}

func (m MTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	now := time.Now()
	for _, middleware := range m.middlewares {
		middleware(request)
	}

	defer func() {
		elapsed := time.Since(now)
		fmt.Println("request took, ", elapsed)
	}()

	return m.rt.RoundTrip(request)
}

func NewHTTPClient(middlewares ...Middleware) *http.Client {
	mt := MTripper{
		rt: http.DefaultTransport,
	}

	for _, middleware := range middlewares {
		mt.Use(middleware)
	}

	c := http.Client{
		Transport: mt,
	}

	return &c
}
