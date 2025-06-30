package main

import (
	"log"
	"net/http"
	"time"
)

func Logger() HandlerFunc {
	return func(ctx *Context) {
		// start Timer
		t := time.Now()
		// process request
		ctx.Next()
		// calculate resolution time
		log.Printf("[%d] %s in %v", ctx.StatusCode, ctx.Req.URL.Path, time.Since(t))
	}
}

func v2Middleware() HandlerFunc {
	return func(ctx *Context) {
		// start Timer
		t := time.Now()
		// process request
		ctx.Fail(http.StatusInternalServerError, "Internal Server Error")
		// calculate resolution time
		log.Printf("[%d] %s in %v for group v2", ctx.StatusCode, ctx.Req.URL.Path, time.Since(t))
	}
}

func main() {
	s := New()
	s.Use(Logger())
	s.GET("/index", func(ctx *Context) {
		ctx.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	v1 := s.Group("/v1")
	{
		v1.GET("/", func(ctx *Context) {
			ctx.HTML(http.StatusOK, "<h1>Hello v1</h1>")
		})
		v1.GET("/hello", func(ctx *Context) {
			ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
		})
	}

	v2 := s.Group("/v2")
	v2.Use(v2Middleware())
	{
		v2.GET("/hello/:name", func(ctx *Context) {
			ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Param("name"), ctx.Path)
		})
		v2.POST("/login", func(ctx *Context) {
			ctx.JSON(http.StatusOK, H{
				"username": ctx.PostForm("username"),
				"password": ctx.PostForm("password"),
			})
		})
	}

	s.Run(":9999")
}
