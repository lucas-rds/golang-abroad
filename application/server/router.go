package server

import (
	"net/http"

	"github.com/go-foward/abroad/utils/compare"
)

// Router is an collection of Routes
type Router struct {
	routes []*Route
}

//NewRouter creates a router
func NewRouter() *Router {
	return &Router{
		routes: make([]*Route, 0),
	}
}

// Route is an representation of a path and a callback
type Route struct {
	path    string
	handler func(w http.ResponseWriter, r *http.Request)
}

// GET adds a get HandleFunc into Router mux
func (router *Router) GET(path string, handler http.HandlerFunc) {
	router.AddRoute(path, handler, []string{http.MethodGet})
}

// POST adds a post HandleFunc into Router mux
func (router *Router) POST(path string, handler http.HandlerFunc) {
	router.AddRoute(path, handler, []string{http.MethodPost})
}

// AddRoute adds a Route into Router mux
func (router *Router) AddRoute(path string, handler http.HandlerFunc, methods []string) {
	route := &Route{path: path, handler: router.acceptMethods(handler, methods)}
	router.routes = append(router.routes, route)
}

func (router *Router) acceptMethods(handler http.HandlerFunc, methods []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if compare.StringsContains(methods, r.Method) {
			handler(w, r)
		}
	}
}
