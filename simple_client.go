// Copyright 2019 潘文斌. All rights reserved.
// license that can be found in the LICENSE file.

package ghttpclient

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Send a Request with GET method
func Get(url string, headers map[string]string) (*http.Response, error) {
	return NewClient().Url(url).Headers(headers).Get()
}

// Send a Request with POST method
func Post(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	return NewClient().Url(url).Headers(headers).Body(body).Post()
}

// Send a Request as a json with POST Method
func PostJson(url string, jsonBytes []byte, headers map[string]string) (*http.Response, error) {
	return NewClient().Url(url).Headers(headers).Body(bytes.NewReader(jsonBytes)).
		Header("Content-Type", "application/json").
		Post()
}

// Send a Request as a form with POST method
func PostForm(url string, data url.Values, headers map[string]string) (*http.Response, error) {
	return NewClient().Url(url).Headers(headers).Body(strings.NewReader(data.Encode())).
		Header("Content-Type", "application/x-www-form-urlencoded").
		Post()
}

// Send a Request with PUT method
func Put(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	return NewClient().Url(url).Headers(headers).Body(body).Put()
}

// Send a Request as a json with PUT method
func PutJson(url string, jsonBytes []byte, headers map[string]string) (*http.Response, error) {
	return NewClient().Url(url).Headers(headers).Body(bytes.NewReader(jsonBytes)).
		Header("Content-Type", "application/json").
		Put()
}

// Send a Request with PATCH method
func Patch(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	return NewClient().Url(url).Headers(headers).Body(body).Patch()
}

// Send a Request with DELETE method
func Delete(url string, headers map[string]string) (*http.Response, error) {
	return NewClient().Url(url).Headers(headers).Delete()
}

// Send a Request with OPTIONS method
func Options(url string, headers map[string]string) (*http.Response, error) {
	return NewClient().Url(url).Headers(headers).Options()
}
