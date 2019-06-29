// Copyright 2019 潘文斌. All rights reserved.
// license that can be found in the LICENSE file.

package ghttpclient

import (
	"io"
	"net/http"
)

func Get(url string, headers map[string]string) (*http.Response, error) {
	return NewClient().Url(url).Headers(headers).Get()
}

func Post(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	return NewClient().Url(url).Headers(headers).Body(body).Post()
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
