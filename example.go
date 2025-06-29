package main

import "net/http"

func main() {
	s := New()
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
