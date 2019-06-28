package ghttpclient

import (
	"io"
	"net/http"
)

func Get(url string, headers map[string]string) (*http.Response, error) {
	return NewClient().SetUrl(url).SetHeaders(headers).Get()
}

func Post(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	return NewClient().SetUrl(url).SetHeaders(headers).SetBody(body).Post()
}

func Put(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	return NewClient().SetUrl(url).SetHeaders(headers).SetBody(body).Put()
}

func Patch(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	return NewClient().SetUrl(url).SetHeaders(headers).SetBody(body).Patch()
}

func Delete(url string, headers map[string]string) (*http.Response, error) {
	return NewClient().SetUrl(url).SetHeaders(headers).Delete()
}

func Options(url string, headers map[string]string) (*http.Response, error) {
	return NewClient().SetUrl(url).SetHeaders(headers).Options()
}
