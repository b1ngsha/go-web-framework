package main

import (
	"fmt"
	"go-web-framework/server"
	"net/http"
)

func main() {
	s := server.New()
	s.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})

	s.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)

		}
	})

	s.Run(":9999")
}
