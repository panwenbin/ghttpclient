// Copyright 2019 潘文斌. All rights reserved.
// license that can be found in the LICENSE file.

package ghttpclient

import (
	"crypto/tls"
	"errors"
	"github.com/panwenbin/ghttpclient/header"
	"io"
	"net/http"
	"time"
)

// GHttpClient is a Method chaining HTTP Client which is based on net/http.Client
// NewClient => set attributes of a request => do the request with an action(Get, Post...)
type GHttpClient struct {
	request       *http.Request
	client        *http.Client
	url           string
	sslSkipVerify bool
	header        header.GHttpHeader
	body          io.Reader
	timeout       time.Duration
}

// NewClient Returns a new GHttpClient
func NewClient() *GHttpClient {
	return &GHttpClient{
		header: make(header.GHttpHeader),
	}
}

// Url sets the url to request to
func (g *GHttpClient) Url(url string) *GHttpClient {
	g.url = url
	return g
}

// Header sets a header
func (g *GHttpClient) Header(headerKey, headerValue string) *GHttpClient {
	g.header.Set(headerKey, headerValue)
	return g
}

// Headers sets a group of headers
func (g *GHttpClient) Headers(httpHeader header.GHttpHeader) *GHttpClient {
	for headerKey, headerValue := range httpHeader {
		g.header.Set(headerKey, headerValue)
	}
	return g
}

// ContentType sets the Content-Type header
func (g *GHttpClient) ContentType(contentType string) *GHttpClient {
	g.header.ContentType(contentType)
	return g
}

// UserAgent sets the User-Agent header
func (g *GHttpClient) UserAgent(userAgent string) *GHttpClient {
	g.header.UserAgent(userAgent)
	return g
}

// Body sets the body of the request
func (g *GHttpClient) Body(body io.Reader) *GHttpClient {
	g.body = body
	return g
}

// SslSkipVerify sets whether or not skipping ssl verify
func (g *GHttpClient) SslSkipVerify(skip bool) *GHttpClient {
	g.sslSkipVerify = skip
	return g
}

// Timeout sets a timeout to the http client
func (g *GHttpClient) Timeout(timeout time.Duration) *GHttpClient {
	g.timeout = timeout
	return g
}

// prepare checks whether attributes are set, and build a http client
func (g *GHttpClient) prepare(method string) error {
	if g.url == "" {
		return errors.New("URL must be set before sending a request")
	}

	request, err := http.NewRequest(method, g.url, g.body)
	if err != nil {
		return err
	}
	request.Header = g.header.ToHttpHeader()

	g.request = request

	g.client = &http.Client{}

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

// send do send the request
func (g *GHttpClient) send() (*http.Response, error) {
	response, err := g.client.Do(g.request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Head sends the Request with HEAD method
func (g *GHttpClient) Head() (*http.Response, error) {
	err := g.prepare("HEAD")
	if err != nil {
		return nil, err
	}

	return g.send()
}

// Get sends the Request with GET method
func (g *GHttpClient) Get() (*http.Response, error) {
	err := g.prepare("GET")
	if err != nil {
		return nil, err
	}

	return g.send()
}

// Post sends the Request with POST method
func (g *GHttpClient) Post() (*http.Response, error) {
	err := g.prepare("POST")
	if err != nil {
		return nil, err
	}

	return g.send()
}

// Put sends the Request with PUT method
func (g *GHttpClient) Put() (*http.Response, error) {
	err := g.prepare("PUT")
	if err != nil {
		return nil, err
	}

	return g.send()
}

// Patch sends the Request with PATCH method
func (g *GHttpClient) Patch() (*http.Response, error) {
	err := g.prepare("PATCH")
	if err != nil {
		return nil, err
	}

	return g.send()
}

// Delete sends the Request with DELETE method
func (g *GHttpClient) Delete() (*http.Response, error) {
	err := g.prepare("DELETE")
	if err != nil {
		return nil, err
	}
	return g.send()
}

// Options sends the Request with OPTIONS method
func (g *GHttpClient) Options() (*http.Response, error) {
	err := g.prepare("OPTIONS")
	if err != nil {
		return nil, err
	}

	return g.send()
}
