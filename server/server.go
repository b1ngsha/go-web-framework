package server

import (
	"fmt"
	"net/http"
)

// HandlerFunc defines the request handler used by server
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// Engine implements the interface of ServeHTTP
type Engine struct {
	router map[string]HandlerFunc
}

// New is the constructor of server.Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// addRoute registers the route
func (engine *Engine) addRoute(method, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// GET defines the method to add GET api
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST api
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
