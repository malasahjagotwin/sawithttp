# sawithttp

**HTTP/1.1 + HTTP/2 client for Go with controllable header order.**

[![Go](https://img.shields.io/badge/go-1.24%2B-00ADD8?style=flat-square&logo=go&logoColor=white)](https://go.dev)
[![License](https://img.shields.io/badge/license-BSD--3--Clause-blue?style=flat-square)](LICENSE)
[![Release](https://img.shields.io/badge/release-v0.1.0-2ea44f?style=flat-square)](https://github.com/malasahjagotwin/sawithttp/tags)
[![Reference](https://img.shields.io/badge/reference-godoc-007d9c?style=flat-square)](https://pkg.go.dev/github.com/malasahjagotwin/sawithttp)

`net/http` is deliberately predictable: it sorts headers, owns the HTTP/2
pseudo-header order and exposes no tuning for frame settings. `sawithttp` gives
you control over all three, which is what makes a request look like a browser
instead of a library.

## Install

```sh
go get github.com/malasahjagotwin/sawithttp@v0.1.0
```

## Quickstart

```go
package main

import (
	"fmt"
	"io"
	"time"

	http "github.com/malasahjagotwin/sawithttp"
	"github.com/malasahjagotwin/sawithttp/http2"
)

func main() {
	tr := &http.Transport{
		ForceAttemptHTTP2:   true,
		MaxIdleConnsPerHost: 8,
		TLSHandshakeTimeout: 8 * time.Second,
	}
	_ = http2.ConfigureTransports(tr)
	client := &http.Client{Transport: tr, Timeout: 15 * time.Second}

	req, _ := http.NewRequest("GET", "https://example.com/", nil)
	req.Header = http.Header{
		http.HeaderOrderKey:  {"user-agent", "accept", "accept-language", "sec-fetch-mode"},
		http.PHeaderOrderKey: {":method", ":authority", ":scheme", ":path"},
		"User-Agent":         {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36"},
		"Accept":             {"text/html,application/xhtml+xml,*/*;q=0.8"},
		"Accept-Language":    {"en-US,en;q=0.9"},
		"Sec-Fetch-Mode":     {"navigate"},
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(resp.Status, len(body))
}
```

## Features

- **Header order** per request (`http.HeaderOrderKey`), not alphabetical.
- **Pseudo-header order** for HTTP/2 (`http.PHeaderOrderKey`).
- **Transport tuning**: connection flow, frame size, priorities, buffer sizes.
- **Drop-in API**: package name is `sawithttp`, so `http.X` call sites keep working.
- **Composable with custom TLS**: bring your own ClientHello/fingerprint layer.

## Packages

| Package | Purpose |
| --- | --- |
| `sawithttp` | HTTP/1.1 client, transport, server |
| `sawithttp/http2` | HTTP/2 transport and server |
| `sawithttp/cookiejar` | RFC 6265 cookie jar |
| `sawithttp/cgi` / `fcgi` | CGI and FastCGI handlers |
| `sawithttp/httputil` | proxy and dump helpers |
| `sawithttp/httptest` | test server helpers |
| `sawithttp/httptrace` | client trace hooks |
| `sawithttp/pprof` | profiling handlers |

## Requirements

Go 1.24 or newer.

## License

BSD-3-Clause. See [LICENSE](LICENSE).
