package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"
)

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := New()
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*.tmpl")
	r.Static("/assets", "static")

	r.GET("/", func(ctx *Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", nil)
	})

	r.Run(":9999")
}
