# Overview
[![Go Reference](https://pkg.go.dev/badge/github.com/gopi-frame/response.svg)](https://pkg.go.dev/github.com/gopi-frame/response)
[]

This is a package for handling HTTP responses.

## Installation

```shell
go get -u -v github.com/gopi-frame/response
```

## Import

```go
import "github.com/gopi-frame/response"
```

### Quick Start

```go
package main

import "net/http"

func main() {
    var handler = func(w http.ResponseWriter, r *http.Request) {
        resp := response.New(http.StatusOK, "Hello World")
        resp.ServeHTTP(w, r)
    }
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

### JSON Response

`JSONResponse` provides a convenient way to send JSON-formatted data as the response body in an HTTP request.

```go
package main

func main() {
    var handler = func(w http.ResponseWriter, r *http.Request) {
        resp := response.New(http.StatusOK).JSON(map[string]string{
            "message": "Hello World",
        })
        resp.ServeHTTP(w, r)
    }
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

## XML Response

`XMLResponse` provides a convenient way to send XML-formatted data as the response body in an HTTP request.

```go
package main

func main() {
    type Response struct {
        Message string `xml:"message"`
    }
    
    var handler = func(w http.ResponseWriter, r *http.Request) {
        resp := response.New(http.StatusOK).XML(Response{
            Message: "Hello World",
        })
        rsp.ServeHTTP(w, r)
    }
    
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

### Redirect
`Redirect` provides a convenient way to redirect an HTTP request.

```go
package main

func main() {
    var handler = func(w http.ResponseWriter, r *http.Request) {
        resp := response.New(http.StatusMovedPermanently).Redirect("https://github.com/gopi-frame/response")
        resp.ServeHTTP(w, r)
    }
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

### Reader Response

`ReaderResponse` provides a convenient way to send the contents of an io.Reader as the response body in an HTTP request.

```go
package main

func main() {
    var handler = func(w http.ResponseWriter, r *http.Request) {
        buffer := bytes.NewBufferString("Hello World")
        resp := response.New(http.StatusOK).Reader(buffer)
        resp.SetContentType("text/plain")
        resp.ServeHTTP(w, r)
    }
    
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

### File Response

`FileResponse` provides a convenient way to send the contents of a file as the response body in an HTTP request.

```go
package main

func main() {
    var handler = func(w http.ResponseWriter, r *http.Request) {
        resp := response.New(http.StatusOK).File("README.md")
        resp.ServeHTTP(w, r)
    }
    
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

### Streamed Response

`StreamedResponse` provides a convenient way to send a streamed response by providing a step function that writes data to the response writer.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    var handler = func(w http.ResponseWriter, r *http.Request) {
        resp := response.New(http.StatusOK).Streamed(func(w http.ResponseWriter) error {
            _, err := w.Write([]byte(fmt.Sprintf("Current time: %v", time.Now())))
            if err != nil {
                return false
            }
            time.Sleep(time.Second)
            return true
        })
        resp.ServeHTTP(w, r)
    }
    
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```