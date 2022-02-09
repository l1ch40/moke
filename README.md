# Moke Web Framework
Moke is a HTTP web framework written in Go (Golang).
It features a use trie-tree parse dynamic route.

## Quick start
```go
package main

import (
	"moke"
	"net/http"
)

func main() {
	r := moke.Default()
	r.GET("/", func(c *moke.Context) {
		c.String(http.StatusOK, "Hello Moke\n")
	})
	r.GET("/hello/:name", func(c *moke.Context) {
		c.JSON(http.StatusOK, moke.H{"message": fmt.Sprintf("Hello %s", c.Param("name"))})
	})
	r.Run(":9999")
}
```
A