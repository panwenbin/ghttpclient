## GHttpClient
A method chaining HTTP Client for Golang with simple methods

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

    defer response.Body.Close()
    body, err := ioutil.ReadAll(response.Body)
    fmt.Printf("%s", body)
}
```

```go
response, err := ghttpclient.NewClient().
    Url("http://www.panwenbin.com/").
    Headers(map[string]string{
        "User-Agent": "ghttpclient",
    }).
    Get()
```