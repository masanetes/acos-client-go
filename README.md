# acos-client-go

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go Reference](https://pkg.go.dev/badge/github.com/masanetes/acos-client-go.svg)](https://pkg.go.dev/github.com/masanetes/acos-client-go)
[![report](https://goreportcard.com/badge/github.com/masanetes/acos-client-go)](https://goreportcard.com/report/github.com/masanetes/acos-client-go)

A simple go client for [a10 networks](https://www.a10networks.com/)' aXAPI

[ACOS 4.1.4-P1 aXAPIv3 Reference Document](https://documentation.a10networks.com/ACOS/414x/ACOS_4_1_4-P1/html/axapiv3/index.html#)

## Documentation

[godoc](https://pkg.go.dev/github.com/masanetes/acos-client-go)

## Supported Version

| ACOS | aXAPI | Status |
|:------|:------|:------|
| 4.1.4 | 3 | :white_check_mark: |

## Install

```
go get -u github.com/masanetes/acos-client-go
```

## Usage

Import

```go
import "github.com/masanetes/acos-client-go"
```

New client

```go
config := client.Config{Host: "<HOST>", User: "<USER>", Pass: "<PASS>", Debug: false}
c, _ := client.New(config)

// Get virtual server
vs, _ := c.Slb.VirtualServer.Get("masanetes-sample-virtualserver")
fmt.Println(vs.Name, vs.IPAddress)
```

[example code](https://github.com/masanetes/acos-client-go/blob/main/example/main.go)

## Optional Parameter

When creating a client, you can accept optional arguments and customize the http client

```go
customHTTPClient := func(client *http.Client) {
	client = &http.Client{} 
}
client.New(config, customHTTPClient)
```

There are also several options available that are likely to be used more frequently and can be used.

**TLS insecure skip verify**

```go
client.New(config, client.InsecureSkipVerify(true))
```

## Authenticate

A token is issued when the client is created, and is added to the header as a bearer token, but if the token expires, the client can re-authenticate.

```go
client.Authenticate()
```

## Handle aXAPI error response

acos-client-go treats HTTP responses of 400 or more from aXAPI as errors. The response is wrapped as an error and can be unwrapped

```go

import "github.com/masanetes/acos-client-go/pkg/axapi/errors"

err = c.Slb.VirtualServer.Delete("not-exist-virtualserver")
if errRes, ok := err.(*errors.ResponseBody); err != nil && ok {
    fmt.Printf("status: %s, msg: %s\n", errRes.Status, errRes.Msg)
}
```

```
status: fail, msg: Object slb virtual-server {not-exist-virtualserver} does not exist
```

## Contribution

I welcome contributions. If you want to fix a problem, make an enhancement, etc., feel free to send a pull request.


