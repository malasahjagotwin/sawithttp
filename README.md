# sawithttp

HTTP client for Go with full control over header order, HTTP/2 pseudo-header
order and connection settings, so requests look like a real browser instead of
a default Go client.

Module path: `sawithttp`

```go
import (
    http "sawithttp"
    "sawithttp/http2"
)
```

## Install

```
require sawithttp v0.0.0
replace sawithttp => ../sawithttp
```

Or pin to the repo:

```
require sawithttp v0.0.0-<timestamp>-<hash>
replace sawithttp => github.com/malasahjagotwin/sawithttp v0.0.0-<timestamp>-<hash>
```

## Usage

```go
package main

import (
    "fmt"
    "io"
    "time"

    http "sawithttp"
    "sawithttp/http2"
)

func main() {
    tr := &http.Transport{
        ForceAttemptHTTP2:   true,
        MaxIdleConnsPerHost: 8,
        TLSHandshakeTimeout: 8 * time.Second,
    }
    _ = http2.ConfigureTransports(tr)
    c := &http.Client{Transport: tr, Timeout: 15 * time.Second}

    req, _ := http.NewRequest("GET", "https://example.com/", nil)
    req.Header = http.Header{
        http.HeaderOrderKey:  {"user-agent", "accept", "accept-language", "sec-fetch-mode"},
        http.PHeaderOrderKey: {":method", ":authority", ":scheme", ":path"},
        "User-Agent":         {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36"},
        "Accept":             {"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"},
        "Accept-Language":    {"en-US,en;q=0.9"},
        "Sec-Fetch-Mode":     {"navigate"},
    }

    resp, err := c.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    body, _ := io.ReadAll(resp.Body)
    fmt.Println(resp.Status, len(body))
}
```

## Features

- Header order and pseudo-header order control (HTTP/1.1 and HTTP/2)
- HTTP/2 settings tuning (connection flow, frame size, priorities)
- TLS layer swap for custom ClientHello fingerprints
- Drop-in `net/http` API (`package sawithttp`), existing code keeps working

## Layout

```
sawithttp        HTTP/1.1 client, transport, server
sawithttp/http2  HTTP/2 transport and server
sawithttp/cgi    CGI handler
sawithttp/fcgi   FastCGI handler
sawithttp/cookiejar
```

## License

BSD-3-Clause, see [LICENSE](LICENSE).
