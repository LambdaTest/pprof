# pprof

[![Run Tests](https://github.com/LambdaTest/pprof/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/LambdaTest/pprof/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/pprof/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/pprof)
[![Go Report Card](https://goreportcard.com/badge/github.com/LambdaTest/pprof)](https://goreportcard.com/report/github.com/LambdaTest/pprof)
[![GoDoc](https://godoc.org/github.com/LambdaTest/pprof?status.svg)](https://godoc.org/github.com/LambdaTest/pprof)
[![Join the chat at https://gitter.im/gin-gonic/gin](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

gin pprof middleware

> Package pprof serves via its HTTP server runtime profiling data in the format expected by the pprof visualization tool.

## Usage

### Start using it

Download and install it:

```bash
go get github.com/LambdaTest/pprof
```

Import it in your code:

```go
import "github.com/LambdaTest/pprof"
```

### Example

```go
package main

import (
  "github.com/LambdaTest/pprof"
  "github.com/LambdaTest/gin"
)

func main() {
  router := gin.Default()
  pprof.Register(router)
  router.Run(":8080")
}
```

### change default path prefix

```go
func main() {
  router := gin.Default()
  // default is "debug/pprof"
  pprof.Register(router, "dev/pprof")
  router.Run(":8080")
}
```

### custom router group

```go
package main

import (
  "net/http"

  "github.com/LambdaTest/pprof"
  "github.com/LambdaTest/gin"
)

func main() {
  router := gin.Default()
  adminGroup := router.Group("/admin", func(c *gin.Context) {
    if c.Request.Header.Get("Authorization") != "foobar" {
      c.AbortWithStatus(http.StatusForbidden)
      return
    }
    c.Next()
  })
  pprof.RouteRegister(adminGroup, "pprof")
  router.Run(":8080")
}

```

### Use the pprof tool

Then use the pprof tool to look at the heap profile:

```bash
go tool pprof http://localhost:8080/debug/pprof/heap
```

Or to look at a 30-second CPU profile:

```bash
go tool pprof http://localhost:8080/debug/pprof/profile
```

Or to look at the goroutine blocking profile, after calling runtime.SetBlockProfileRate in your program:

```bash
go tool pprof http://localhost:8080/debug/pprof/block
```

Or to collect a 5-second execution trace:

```bash
wget http://localhost:8080/debug/pprof/trace?seconds=5
```
