package elastic

import (
	"io"
	"net/http"
)

type elasticHttp struct {
	EndPoint string
	Index    string
	Type     string
	ID       string
	Path     string
}

const typeJSON = "application/json"

func (e *elasticHttp) POST(body io.Reader) (*http.Response, error) {
	return http.Post(e.EndPoint+e.Path, typeJSON, body)
}

func (e *elasticHttp) PUT(body io.Reader) (*http.Response, error) {
	url := e.EndPoint + e.Path
	return e.Request(http.MethodPut, url, "", body)
}

func (e *elasticHttp) GET() (*http.Response, error) {
	url := e.EndPoint + e.Path
	return e.Request(http.MethodGet, url, "", nil)
}

func (e *elasticHttp) Request(medthod, url, contentType string, body io.Reader) (*http.Response, error) {
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
