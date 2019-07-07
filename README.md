## GHttpClient
A method chaining HTTP Client for Golang with simple methods

## Simple Methods
Get, Post, PostJson, PostForm, Put, PutJson, Patch, Delete, Options

## Examples
```go
package main

import (
	"fmt"
	"github.com/panwenbin/ghttpclient"
	"io/ioutil"
	"log"
)

func main() {
    response, err := ghttpclient.Get("http://www.panwenbin.com/", nil)
    if err != nil {
        log.Fatal(err)
    }

    body, err := ghttpclient.ReadBodyClose(response)
    fmt.Printf("%s", body)
}
```

```go
headers := header.GHttpHeader{}
headers.UserAgent("ghttpclient")
response, err := ghttpclient.NewClient().
    Url("http://www.panwenbin.com/").
    Headers(headers).
    Get()
```