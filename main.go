package main

import (
	"Gee/gee"
	"log"
	"net/http"
)

func middle() gee.HandlerFunc {
	return func(c *gee.Context) {
		log.Println("中间件启动")
		c.Next()
	}
}

func main() {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello Go")
	})
	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"Username": c.PostForm("username"),
		})
	})
	r.GET("/:name", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello %s", c.Param("name"), c.Path)
	})
	v1 := r.Group("/v1")
	v1.Use(middle())
	{
		v1.GET("/hello", func(c *gee.Context) {
			c.String(http.StatusOK, "Hello v1")
		})
	}

	r.Run(":8080")
}
