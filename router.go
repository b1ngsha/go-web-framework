package main

import (
	"net/http"
	"strings"
)

// roots key will be like: "GET", "POST", etc
// handlers key will be like: "GET-/p/:lang/doc", "POST-/p/book"
type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

// create new router
func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// parse pattern into parts
func parsePattern(pattern string) []string {
	segments := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, item := range segments {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' { // only one '*' is allowed
				break
			}
		}
	}
	return parts
}

// add route (build the trie tree)
func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

// get route and params map by path
// the params map stores the user input path -> route mapping
// For example:
//  1. user input path: /p/go/doc
//     route: /p/:lang/doc
//     the params map is: {lang: go}
//  2. user input path: /p/go/doc
//     route: /p/*path
//     the params map is: {path: go/doc}
func (r *router) getRoute(method, path string) (*node, map[string]string) {
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	segments := parsePattern(path)
	params := make(map[string]string)
	n := root.search(segments, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for i, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = segments[i]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(segments[i:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

// route handler
func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	c.Next()
}
