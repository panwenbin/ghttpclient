// Copyright 2019 潘文斌. All rights reserved.
// license that can be found in the LICENSE file.

package ghttpclient

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
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

// ReadBodyClose fetches the response Body, then close the Body
// supports gzip content-type
func ReadBodyClose(response *http.Response) ([]byte, error) {
	if response == nil {
		return nil, errors.New("response  is nil")
	}

	defer response.Body.Close()

	var body []byte
	var err error

	if response.Header.Get("Content-Encoding") == "gzip" {
		gzReader, err := gzip.NewReader(response.Body)
		if err != nil {
			return nil, err
		}
		body, err = ioutil.ReadAll(gzReader)
	} else {
		body, err = ioutil.ReadAll(response.Body)
	}
	if err != nil {
		return nil, err
	}

	return body, nil
}

// ReadJsonClose fetches the response Body and try to decode as a json, then close the Body
func ReadJsonClose(response *http.Response, v interface{}) error {
	contentType := response.Header.Get("Content-Type")
	if contentType != "application/json" {
		return errors.New(fmt.Sprintf("content type application/json expected, but %s got", contentType))
	}
	body, err := ReadBodyClose(response)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}

	return nil
}
