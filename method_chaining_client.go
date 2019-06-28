package ghttpclient

import (
	"crypto/tls"
	"errors"
	"io"
	"net/http"
	"time"
)

// Method chaining HTTP Client
type GHttpClient struct {
	request       *http.Request
	client        *http.Client
	url           string
	sslSkipVerify bool
	header        http.Header
	body 		  io.Reader
	timeout       time.Duration
}

func NewClient() *GHttpClient {
	return &GHttpClient{
	}
}

func (g *GHttpClient) Url(url string) *GHttpClient {
	g.url = url
	return g
}

func (g *GHttpClient) Header(headerKey, headerValue string) *GHttpClient {
	g.header.Set(headerKey, headerValue)
	return g
}

func (g *GHttpClient) Headers(headers map[string]string) *GHttpClient {
	for headerKey, headerValue := range headers {
		g.header.Set(headerKey, headerValue)
	}
	return g
}

func (g *GHttpClient) Body(body io.Reader) *GHttpClient {
	g.body = body
	return g
}

func (g *GHttpClient) SslSkipVerify(skip bool) *GHttpClient {
	g.sslSkipVerify = skip
	return g
}

func (g *GHttpClient) Timeout(timeout time.Duration) *GHttpClient {
	g.timeout = timeout
	return g
}

func (g *GHttpClient) prepare(method string) error {
	if g.url == "" {
		return errors.New("URL must be set before sending a request")
	}

	request, err := http.NewRequest(method, g.url, g.body)
	if err != nil {
		return err
	}
	request.Header = g.header

	g.request = request

	g.client = &http.Client{
	}

	if g.timeout > 0 {
		g.client.Timeout = g.timeout
	}

	if g.sslSkipVerify {
		g.client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	return nil
}

func (g *GHttpClient) send() (*http.Response, error) {
	response, err := g.client.Do(g.request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (g *GHttpClient) Head() (*http.Response, error) {
	err := g.prepare("HEAD")
	if err != nil {
		return nil, err
	}

	return g.send()
}

func (g *GHttpClient) Get() (*http.Response, error) {
	err := g.prepare("GET")
	if err != nil {
		return nil, err
	}

	return g.send()
}

func (g *GHttpClient) Post() (*http.Response, error) {
	err := g.prepare("POST")
	if err != nil {
		return nil, err
	}

	return g.send()
}

func (g *GHttpClient) Put() (*http.Response, error) {
	err := g.prepare("PUT")
	if err != nil {
		return nil, err
	}

	return g.send()
}

func (g *GHttpClient) Patch() (*http.Response, error) {
	err := g.prepare("PATCH")
	if err != nil {
		return nil, err
	}

	return g.send()
}

func (g *GHttpClient) Delete() (*http.Response, error) {
	err := g.prepare("DELETE")
	if err != nil {
		return nil, err
	}
	return g.send()
}

func (g *GHttpClient) Options() (*http.Response, error) {
	err := g.prepare("OPTIONS")
	if err != nil {
		return nil, err
	}

	return g.send()
}