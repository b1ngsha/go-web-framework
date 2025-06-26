package server

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

// create new router
func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

// add route
func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
	log.Printf("Route %s - %s\n", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

// route handler
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
