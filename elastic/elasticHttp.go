package elastic

import (
	"io"
	"net/http"
)

type elasticHTTP struct {
	EndPoint string
	Index    string
	Type     string
	ID       string
	Path     string
}

const typeJSON = "application/json"

func (e *elasticHTTP) POST(body io.Reader) (*http.Response, error) {
	return http.Post(e.EndPoint+e.Path, typeJSON, body)
}

func (e *elasticHTTP) PUT(body io.Reader) (*http.Response, error) {
	url := e.EndPoint + e.Path
	return e.Request(http.MethodPut, url, "", body)
}

func (e *elasticHTTP) GET() (*http.Response, error) {
	url := e.EndPoint + e.Path
	return e.Request(http.MethodGet, url, "", nil)
}

func (e *elasticHTTP) Request(medthod, url, contentType string, body io.Reader) (*http.Response, error) {
	c := http.DefaultClient
	req, err := http.NewRequest(medthod, url, body)

	if err != nil {
		return nil, err
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	return c.Do(req)
}
