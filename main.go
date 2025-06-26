package main

import (
	"go-web-framework/server"
	"net/http"
)

func main() {
	s := server.New()
	s.GET("/", func(ctx *server.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello World!</h1>")
	})

	s.GET("/hello", func(ctx *server.Context) {
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
	})

	s.POST("/login", func(ctx *server.Context) {
		ctx.JSON(http.StatusOK, server.H{
			"username": ctx.PostForm("username"),
			"password": ctx.PostForm("password"),
		})
	})

	s.Run(":9999")
}
