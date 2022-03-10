package main

import (
	"Gee/gee"
	"net/http"
)

func main() {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello Go\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"geek"}
		c.String(http.StatusOK, names[4])
	})

	r.Run(":9999")
}
