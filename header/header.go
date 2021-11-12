// Copyright 2019 潘文斌. All rights reserved.
// license that can be found in the LICENSE file.

package header

import "net/http"

// GHttpHeader is the header struct for GHttpClient
// To init a GHttpHeader, use make(GHttpHeader)
type GHttpHeader map[string]string

// Set sets a pare of header key and value
func (h GHttpHeader) Set(key, value string) GHttpHeader {
	h[key] = value
	return h
}

// Del deletes a header by key
func (h GHttpHeader) Del(key string) GHttpHeader {
	delete(h, key)
	return h
}

// ContentType sets the Content-Type header
func (h GHttpHeader) ContentType(contentType string) GHttpHeader {
	return h.Set("Content-Type", contentType)
}

// UserAgent sets the User-Agent header
func (h GHttpHeader) UserAgent(userAgent string) GHttpHeader {
	return h.Set("User-Agent", userAgent)
}

// AcceptEncodingGzip sets the Accept-Encoding to gzip
func (h GHttpHeader) AcceptEncodingGzip() GHttpHeader {
	return h.Set("Accept-Encoding", "gzip")
}

// ContentEncodingZip sets the Content-Encoding to gzip
func (h GHttpHeader) ContentEncodingZip() GHttpHeader {
	return h.Set("Content-Encoding", "gzip")
}

// RemoveContentEncoding deletes the Content-Encoding
func (h GHttpHeader) RemoveContentEncoding() GHttpHeader {
	return h.Del("Content-Encoding")
}

// IsContentEncodingZip checks if Content-Encoding is gzip
func (h GHttpHeader) IsContentEncodingZip() bool {
	v, _ := h["Content-Encoding"]
	return v == "gzip"
}

// ToHttpHeader converts GHttpHeader to http.Header
func (h GHttpHeader) ToHttpHeader() http.Header {
	httpHeader := make(http.Header)
	for headerKey, headerValue := range h {
		httpHeader.Set(headerKey, headerValue)
	}
	return httpHeader
}
