// Copyright 2019 潘文斌. All rights reserved.
// license that can be found in the LICENSE file.

package ghttpclient_test

import (
	"bytes"
	"encoding/json"
	"github.com/panwenbin/ghttpclient"
	"github.com/panwenbin/ghttpclient/header"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"math/rand"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	body, err := ghttpclient.Get("http://cp.fei.lv/ghttpclient", nil).ReadBodyClose()
	if err != nil {
		t.Error("error occurs")
		t.Fatal(err)
	}

	if strings.Compare("ghttpclient", string(body)) != 0 {
		t.Fatalf("expect 'ghttpclient, got %s", body)
	}
}

func TestPost(t *testing.T) {
	body, err := ghttpclient.Post("http://cp.fei.lv/", strings.NewReader("ghttpclient"), nil).ReadBodyClose()
	if err != nil {
		t.Error("error occurs")
		t.Fatal(err)
	}

	if strings.Compare("ghttpclient", string(body)) != 0 {
		t.Fatalf("expect 'ghttpclient, got %s", body)
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

	bodyJsonBytes, err := ghttpclient.PostJson("http://cp.fei.lv/", jsonBytes, nil).ReadBodyClose()
	if err != nil {
		t.Error("error occurs")
		t.Fatal(err)
	}

	if bytes.Compare(jsonBytes, bodyJsonBytes) != 0 {
		t.Errorf("expect '%s', got %s", jsonBytes, bodyJsonBytes)
	}
}

func TestPostForm(t *testing.T) {
	data := url.Values{}
	data.Add("msg", "ghttpclient")

	body, err := ghttpclient.PostForm("http://cp.fei.lv/", data, nil).ReadBodyClose()
	if err != nil {
		t.Error("error occurs")
		t.Fatal(err)
	}

	if strings.Compare(data.Encode(), string(body)) != 0 {
		t.Errorf("expect 'ghttpclient, got %s", body)
	}
}

func TestPut(t *testing.T) {
	body, err := ghttpclient.Put("http://cp.fei.lv/put", strings.NewReader("ghttpclient"), nil).ReadBodyClose()
	if err != nil {
		t.Error("error occurs")
		t.Fatal(err)
	}

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

	var bodyJson TestJson
	err := ghttpclient.PutJson("http://cp.fei.lv/ghttpclient", jsonBytes, nil).ReadJsonClose(&bodyJson)
	if err != nil {
		t.Error("error occurs")
		t.Fatal(err)
	}

	bodyJsonBytes, _ := json.Marshal(bodyJson)
	if bytes.Compare(jsonBytes, bodyJsonBytes) != 0 {
		t.Errorf("expect '%s', got %s", jsonBytes, bodyJsonBytes)
	}
}

func TestPatch(t *testing.T) {
	body, err := ghttpclient.Patch("http://cp.fei.lv/patch", strings.NewReader("ghttpclient"), nil).ReadBodyClose()
	if err != nil {
		t.Error("error occurs")
		t.Fatal(err)
	}

	if strings.Compare("ghttpclient", string(body)) != 0 {
		t.Errorf("expect 'ghttpclient, got %s", body)
	}
}

func TestDelete(t *testing.T) {
	body, err := ghttpclient.Delete("http://cp.fei.lv/ghttpclient", nil).ReadBodyClose()
	if err != nil {
		t.Error("error occurs")
		t.Fatal(err)
	}

	if strings.Compare("ghttpclient", string(body)) != 0 {
		t.Errorf("expect 'ghttpclient, got %s", body)
	}
}

func TestOptions(t *testing.T) {
	body, err := ghttpclient.Options("http://cp.fei.lv/ghttpclient", nil).ReadBodyClose()
	if err != nil {
		t.Error("error occurs")
		t.Fatal(err)
	}

	if strings.Compare("ghttpclient", string(body)) != 0 {
		t.Errorf("expect 'ghttpclient, got %s", body)
	}
}

func TestGetWithHeader(t *testing.T) {
	headers := header.GHttpHeader{}
	headers.UserAgent("ghttpclient")
	body, err := ghttpclient.Get("http://ua.fei.lv/", headers).ReadBodyClose()

	if err != nil {
		t.Error("error occurs")
		t.Fatal(err)
	}

	if strings.Compare("ghttpclient", string(body)) != 0 {
		t.Fatalf("expect 'ghttpclient, got %s", body)
	}
}

func TestGetWithGzip(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	buffer := bytes.Buffer{}
	for i := 0; i < 1024; i++ {
		buffer.WriteRune(rune(r.Intn(26) + 65))
	}
	headers := header.GHttpHeader{}
	headers.AcceptEncodingGzip()
	body, err := ghttpclient.Post("http://cp.fei.lv/ghttpclient", bytes.NewReader(buffer.Bytes()), headers).ReadBodyClose()
	if err != nil {
		t.Error("error occurs")
		t.Fatal(err)
	}

	if bytes.Compare(buffer.Bytes(), body) != 0 {
		t.Fatalf("expect '%s', got %s", buffer.Bytes(), body)
	}
}

func TestTryUTF8ReadBodyClose(t *testing.T) {
	utf8Str := "简体中文"
	gbkStr, _, _ := transform.String(simplifiedchinese.GBK.NewEncoder(), utf8Str)

	body, err := ghttpclient.Post("http://cp.fei.lv/gbk", strings.NewReader(gbkStr), nil).TryUTF8ReadBodyClose()
	if err != nil {
		t.Error("error occurs")
		t.Fatal(err)
	}

	if strings.Compare(utf8Str, string(body)) != 0 {
		t.Fatalf("expect '%s', got %s", utf8Str, body)
	}
}
