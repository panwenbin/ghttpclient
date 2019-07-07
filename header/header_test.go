package header_test

import (
	"github.com/panwenbin/ghttpclient/header"
	"testing"
)

func TestGHttpHeader_Set(t *testing.T) {
	headers := header.GHttpHeader{}
	headers.Set("Content-Type", header.CONTENT_TYPE_JSON)
	contentType := headers.ToHttpHeader().Get("Content-Type")
	if contentType != header.CONTENT_TYPE_JSON {
		t.Errorf("expect %s, got %s", header.CONTENT_TYPE_JSON, contentType)
	}
}

func TestGHttpHeader_ContentType(t *testing.T) {
	headers := header.GHttpHeader{}
	headers.ContentType(header.CONTENT_TYPE_JSON)
	contentType := headers.ToHttpHeader().Get("Content-Type")
	if contentType != header.CONTENT_TYPE_JSON {
		t.Errorf("expect %s, got %s", header.CONTENT_TYPE_JSON, contentType)
	}
}

func TestGHttpHeader_UserAgent(t *testing.T) {
	headers := header.GHttpHeader{}
	headers.UserAgent("ghttpclient")
	userAgent := headers.ToHttpHeader().Get("User-Agent")
	if userAgent != "ghttpclient" {
		t.Errorf("expect ghttpclient, got %s", userAgent)
	}
}
