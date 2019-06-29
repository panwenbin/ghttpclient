// Copyright 2019 潘文斌. All rights reserved.
// license that can be found in the LICENSE file.

package ghttpclient_test

import (
	"github.com/panwenbin/ghttpclient"
	"io/ioutil"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {
	response, err := ghttpclient.Get("http://cp.fei.lv/ghttpclient", nil)
	if err != nil {
		t.Error("network error")
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if strings.Compare("ghttpclient", string(body)) != 0 {
		t.Errorf("expect 'ghttpclient, got %s", body)
	}
}

func TestPost(t *testing.T) {
	response, err := ghttpclient.Post("http://cp.fei.lv/", strings.NewReader("ghttpclient"), nil)
	if err != nil {
		t.Error("network error")
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if strings.Compare("ghttpclient", string(body)) != 0 {
		t.Errorf("expect 'ghttpclient, got %s", body)
	}
}

func TestPut(t *testing.T) {
	response, err := ghttpclient.Put("http://cp.fei.lv/put", strings.NewReader("ghttpclient"), nil)
	if err != nil {
		t.Error("network error")
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if strings.Compare("ghttpclient", string(body)) != 0 {
		t.Errorf("expect 'ghttpclient, got %s", body)
	}
}

func TestPatch(t *testing.T) {
	response, err := ghttpclient.Patch("http://cp.fei.lv/patch", strings.NewReader("ghttpclient"), nil)
	if err != nil {
		t.Error("network error")
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if strings.Compare("ghttpclient", string(body)) != 0 {
		t.Errorf("expect 'ghttpclient, got %s", body)
	}
}

func TestDelete(t *testing.T) {
	response, err := ghttpclient.Delete("http://cp.fei.lv/ghttpclient", nil)
	if err != nil {
		t.Error("network error")
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if strings.Compare("ghttpclient", string(body)) != 0 {
		t.Errorf("expect 'ghttpclient, got %s", body)
	}
}

func TestOptions(t *testing.T) {
	response, err := ghttpclient.Options("http://cp.fei.lv/ghttpclient", nil)
	if err != nil {
		t.Error("network error")
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if strings.Compare("ghttpclient", string(body)) != 0 {
		t.Errorf("expect 'ghttpclient, got %s", body)
	}
}

func TestGetWithHeader(t *testing.T) {
	response, err := ghttpclient.Get("http://ua.fei.lv/", map[string]string {
		"User-Agent": "ghttpclient",
	})

	if err != nil {
		t.Error("network error")
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if strings.Compare("ghttpclient", string(body)) != 0 {
		t.Errorf("expect 'ghttpclient, got %s", body)
	}
}