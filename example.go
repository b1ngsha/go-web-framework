package main

import "net/http"

func main() {
	s := New()
	s.GET("/", func(ctx *Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello World!</h1>")
	})

	s.GET("/hello", func(ctx *Context) {
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
	})

	s.GET("/hello/:name", func(ctx *Context) {
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Param("name"), ctx.Path)
	})

	s.GET("/assets/*filepath", func(ctx *Context) {
		ctx.JSON(http.StatusOK, H{
			"filepath": ctx.Param("filepath"),
		})
	})

	s.Run(":9999")
}
