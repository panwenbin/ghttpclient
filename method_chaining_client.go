// Copyright 2019 潘文斌. All rights reserved.
// license that can be found in the LICENSE file.

package ghttpclient

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/panwenbin/ghttpclient/header"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

var (
	green  = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white  = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	red    = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	reset  = string([]byte{27, 91, 48, 109})
)

// GHttpClient is a Method chaining HTTP Client which is based on net/http.Client
// NewClient => set attributes of a request => do the request with an action(Get, Post...)
type GHttpClient struct {
	request       *http.Request
	client        *http.Client
	url           string
	sslSkipVerify bool
	noRedirect    bool
	header        header.GHttpHeader
	body          io.Reader
	cookieJar     http.CookieJar
	timeout       time.Duration
	response      *http.Response
	err           error
	debug         bool
	startTime     time.Time
	logger        *log.Logger
}

// NewClient Returns a new GHttpClient
func NewClient() *GHttpClient {
	return &GHttpClient{
		header: make(header.GHttpHeader),
		logger: log.New(os.Stdout, "ghttpclient", log.Ldate|log.Ltime|log.Lmicroseconds),
		debug:  Debug,
	}
}

// DebugOn sets debug to true
func (g *GHttpClient) DebugOn() *GHttpClient {
	g.debug = true
	return g
}

// Debug sets debug to the value of param
func (g *GHttpClient) Debug(debug bool) *GHttpClient {
	g.debug = debug
	return g
}

func (g *GHttpClient) LogDebug(flag string) {
	now := time.Now()
	var content string
	if flag == "S" {
		content = ""
	} else {
		statusCode := 0
		if g.response != nil {
			statusCode = g.response.StatusCode
		}
		content = fmt.Sprintf(" %s[%d]%s %3.3fs", statusCodeColor(statusCode), statusCode, reset, now.Sub(g.startTime).Seconds())
	}

	str := fmt.Sprintf("[GHTTP] %s [%3s] [%s]%s %s\r\n",
		now.Format("2006-01-02 15:04:05.000"),
		g.request.Method,
		flag,
		content,
		g.request.URL)
	g.logger.Writer().Write([]byte(str))
}

// statusCodeColor returns a color for displaying in terminal.
func statusCodeColor(code int) string {
	switch {
	case code < http.StatusContinue:
		return red
	case code >= http.StatusContinue && code < http.StatusMultipleChoices:
		return green
	case code >= http.StatusMultipleChoices && code < http.StatusBadRequest:
		return white
	case code >= http.StatusBadRequest && code < http.StatusInternalServerError:
		return yellow
	default:
		return red
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

// CookieJar sets the cookie jar of the request
func (g *GHttpClient) CookieJar(cookieJar http.CookieJar) *GHttpClient {
	g.cookieJar = cookieJar
	return g
}

// SslSkipVerify sets whether or not skipping ssl verify
func (g *GHttpClient) SslSkipVerify(skip bool) *GHttpClient {
	g.sslSkipVerify = skip
	return g
}

// NoRedirect sets whether or not to stop following redirects
func (g *GHttpClient) NoRedirect(noFollow bool) *GHttpClient {
	g.noRedirect = noFollow
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

	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	if g.cookieJar != nil {
		g.client.Jar = g.cookieJar
	}

	if g.timeout > 0 {
		g.client.Timeout = g.timeout
	}

	if g.sslSkipVerify {
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	if g.noRedirect {
		g.client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	g.client.Transport = transport

	return nil
}

// send do send the request
func (g *GHttpClient) send() *GHttpClient {
	if g.debug {
		g.startTime = time.Now()
		g.LogDebug("S")
	}
	g.response, g.err = g.client.Do(g.request)
	if g.debug {
		g.LogDebug("R")
	}

	return g
}

// Head sends the Request with HEAD method
func (g *GHttpClient) Head() *GHttpClient {
	g.err = g.prepare("HEAD")
	if g.err != nil {
		return g
	}

	return g.send()
}

// Get sends the Request with GET method
func (g *GHttpClient) Get() *GHttpClient {
	g.err = g.prepare("GET")
	if g.err != nil {
		return g
	}

	return g.send()
}

// Post sends the Request with POST method
func (g *GHttpClient) Post() *GHttpClient {
	g.err = g.prepare("POST")
	if g.err != nil {
		return g
	}

	return g.send()
}

// Put sends the Request with PUT method
func (g *GHttpClient) Put() *GHttpClient {
	g.err = g.prepare("PUT")
	if g.err != nil {
		return g
	}

	return g.send()
}

// Patch sends the Request with PATCH method
func (g *GHttpClient) Patch() *GHttpClient {
	g.err = g.prepare("PATCH")
	if g.err != nil {
		return g
	}

	return g.send()
}

// Delete sends the Request with DELETE method
func (g *GHttpClient) Delete() *GHttpClient {
	g.err = g.prepare("DELETE")
	if g.err != nil {
		return g
	}
	return g.send()
}

// Options sends the Request with OPTIONS method
func (g *GHttpClient) Options() *GHttpClient {
	g.err = g.prepare("OPTIONS")
	if g.err != nil {
		return g
	}

	return g.send()
}

// Response returns http.Response and error
func (g *GHttpClient) Response() (*http.Response, error) {
	return g.response, g.err
}

// ReadBodyClose fetches the response Body, then close the Body
// supports gzip content-type
func (g *GHttpClient) ReadBodyClose() ([]byte, error) {
	if g.debug {
		defer g.LogDebug("E")
	}
	if g.err != nil {
		return []byte{}, g.err
	}
	return ReadBodyClose(g.response)
}

// TryUTF8ReadBodyClose tries to transfer the body bytes to utf-8 bytes when the body bytes is not in utf-8 encoding
func (g *GHttpClient) TryUTF8ReadBodyClose() ([]byte, error) {
	if g.debug {
		defer g.LogDebug("E")
	}
	if g.err != nil {
		return []byte{}, g.err
	}
	return TryUTF8ReadBodyClose(g.response)
}

// ReadJsonClose fetches the response Body and try to decode as a json, then close the Body
func (g *GHttpClient) ReadJsonClose(v interface{}) error {
	if g.debug {
		defer g.LogDebug("E")
	}
	if g.err != nil {
		return g.err
	}
	return ReadJsonClose(g.response, v)
}
