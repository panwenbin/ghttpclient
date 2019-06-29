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

func Get(url string, headers map[string]string) (*http.Response, error) {
	return NewClient().Url(url).Headers(headers).Get()
}

func Post(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	return NewClient().Url(url).Headers(headers).Body(body).Post()
}

func PostJson(url string, jsonBytes []byte, headers map[string]string) (*http.Response, error) {
	return NewClient().Url(url). Headers(headers). Body(bytes.NewReader(jsonBytes)).
		Header("Content-Type", "application/json").
		Post()
}

func PostForm(url string, data url.Values, headers map[string]string) (*http.Response, error) {
	return NewClient().Url(url). Headers(headers). Body(strings.NewReader(data.Encode())).
		Header("Content-Type", "application/x-www-form-urlencoded").
		Post()
}

func Put(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	return NewClient().Url(url).Headers(headers).Body(body).Put()
}

func Patch(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	return NewClient().Url(url).Headers(headers).Body(body).Patch()
}

func Delete(url string, headers map[string]string) (*http.Response, error) {
	return NewClient().Url(url).Headers(headers).Delete()
}

func Options(url string, headers map[string]string) (*http.Response, error) {
	return NewClient().Url(url).Headers(headers).Options()
}
