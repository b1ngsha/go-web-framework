package server

import (
	"net/http"
)

// HandlerFunc defines the request handler used by server
type HandlerFunc func(*Context)

// Engine implements the interface of ServeHTTP
type Engine struct {
	router *router
}

// New is the constructor of server.Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

// addRoute registers the route
func (engine *Engine) addRoute(method, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
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

// ServeHTTP implements the interface of http.Handler
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
