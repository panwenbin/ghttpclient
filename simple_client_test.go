// Copyright 2019 潘文斌. All rights reserved.
// license that can be found in the LICENSE file.

package ghttpclient_test

import (
	"bytes"
	"encoding/json"
	"github.com/panwenbin/ghttpclient"
	"github.com/panwenbin/ghttpclient/header"
	"net/url"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {
	response, err := ghttpclient.Get("http://cp.fei.lv/ghttpclient", nil)
	if err != nil {
		t.Error("network error")
	}

	body, err := ghttpclient.ReadBodyClose(response)
	if strings.Compare("ghttpclient", string(body)) != 0 {
		t.Errorf("expect 'ghttpclient, got %s", body)
	}
}

func TestPost(t *testing.T) {
	response, err := ghttpclient.Post("http://cp.fei.lv/", strings.NewReader("ghttpclient"), nil)
	if err != nil {
		t.Error("network error")
	}

	body, err := ghttpclient.ReadBodyClose(response)
	if strings.Compare("ghttpclient", string(body)) != 0 {
		t.Errorf("expect 'ghttpclient, got %s", body)
	}
}

func TestPostJson(t *testing.T) {
	type TestJson struct {
		Msg string `json:"msg"`
	}
	testJson := TestJson{
		Msg: "ghttpclient",
	}
	jsonBytes, _ := json.Marshal(testJson)

	response, err := ghttpclient.PostJson("http://cp.fei.lv/", jsonBytes, nil)
	if err != nil {
		t.Error("network error")
	}

	bodyJsonBytes, err := ghttpclient.ReadBodyClose(response)
	if bytes.Compare(jsonBytes, bodyJsonBytes) != 0 {
		t.Errorf("expect 'ghttpclient, got %s", bodyJsonBytes)
	}
}

func TestPostForm(t *testing.T) {
	data := url.Values{}
	data.Add("msg", "ghttpclient")

	response, err := ghttpclient.PostForm("http://cp.fei.lv/", data, nil)
	if err != nil {
		t.Error("network error")
	}

	body, err := ghttpclient.ReadBodyClose(response)
	if strings.Compare(data.Encode(), string(body)) != 0 {
		t.Errorf("expect 'ghttpclient, got %s", body)
	}
}

func TestPut(t *testing.T) {
	response, err := ghttpclient.Put("http://cp.fei.lv/put", strings.NewReader("ghttpclient"), nil)
	if err != nil {
		t.Error("network error")
	}

	body, err := ghttpclient.ReadBodyClose(response)
	if strings.Compare("ghttpclient", string(body)) != 0 {
		t.Errorf("expect 'ghttpclient, got %s", body)
	}
}

func TestPutJson(t *testing.T) {
	type TestJson struct {
		Msg string `json:"msg"`
	}
	testJson := TestJson{
		Msg: "ghttpclient",
	}
	jsonBytes, _ := json.Marshal(testJson)

	response, err := ghttpclient.PutJson("http://cp.fei.lv/ghttpclient", jsonBytes, nil)
	if err != nil {
		t.Error("network error")
	}

	body, err := ghttpclient.ReadBodyClose(response)
	if bytes.Compare(jsonBytes, body) != 0 {
		t.Errorf("expect 'ghttpclient, got %s", body)
	}
}

func TestPatch(t *testing.T) {
	response, err := ghttpclient.Patch("http://cp.fei.lv/patch", strings.NewReader("ghttpclient"), nil)
	if err != nil {
		t.Error("network error")
	}

	body, err := ghttpclient.ReadBodyClose(response)
	if strings.Compare("ghttpclient", string(body)) != 0 {
		t.Errorf("expect 'ghttpclient, got %s", body)
	}
}

func TestDelete(t *testing.T) {
	response, err := ghttpclient.Delete("http://cp.fei.lv/ghttpclient", nil)
	if err != nil {
		t.Error("network error")
	}

	body, err := ghttpclient.ReadBodyClose(response)
	if strings.Compare("ghttpclient", string(body)) != 0 {
		t.Errorf("expect 'ghttpclient, got %s", body)
	}
}

func TestOptions(t *testing.T) {
	response, err := ghttpclient.Options("http://cp.fei.lv/ghttpclient", nil)
	if err != nil {
		t.Error("network error")
	}

	body, err := ghttpclient.ReadBodyClose(response)
	if strings.Compare("ghttpclient", string(body)) != 0 {
		t.Errorf("expect 'ghttpclient, got %s", body)
	}
}

func TestGetWithHeader(t *testing.T) {
	headers := header.GHttpHeader{}
	headers.UserAgent("ghttpclient")
	response, err := ghttpclient.Get("http://ua.fei.lv/", headers)

	if err != nil {
		t.Error("network error")
	}

	body, err := ghttpclient.ReadBodyClose(response)
	if strings.Compare("ghttpclient", string(body)) != 0 {
		t.Errorf("expect 'ghttpclient, got %s", body)
	}
}
