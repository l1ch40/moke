package main

import (
	"fmt"
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
	r.GET("/panic", func(c *moke.Context) {
		names := []string{"gogogo"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
