package main

import (
	"abroad/utils/compare"
	"net/http"
)

/*
	Trying to encapsulate and add functionality into default golang server
	Also trying to keep it simple.
*/

// Middleware is a function that mas execute next to continue the request flow
type Middleware func(next http.HandlerFunc) http.HandlerFunc

// Route is an representation of a path and a callback
type Route struct {
	path    string
	handler func(w http.ResponseWriter, r *http.Request)
}

// Server is an net/http server with
type Server struct {
	mux         *http.ServeMux
	routes      []*Route
	middlewares []Middleware
}

// NewServer is a Server builder
func NewServer() Server {
	return Server{
		mux:         http.NewServeMux(),
		routes:      make([]*Route, 0),
		middlewares: make([]Middleware, 0),
	}
}

// GET adds a get HandleFunc into server mux
func (server *Server) GET(path string, callback func(w http.ResponseWriter, r *http.Request)) {
	server.AddRoute(path, callback, []string{http.MethodGet})
}

// POST adds a post HandleFunc into server mux
func (server *Server) POST(path string, callback func(w http.ResponseWriter, r *http.Request)) {
	server.AddRoute(path, callback, []string{http.MethodPost})
}

// AddRoute adds a Route into server mux
func (server *Server) AddRoute(path string, callback func(w http.ResponseWriter, r *http.Request), methods []string) {
	handler := &Route{path: path, handler: server.acceptMethods(callback, methods)}
	server.routes = append(server.routes, handler)
}

// AddMiddleware adds a Middleware into server to decorate all routes
func (server *Server) AddMiddleware(callback Middleware) {
	server.middlewares = append(server.middlewares, callback)
}

func (server *Server) registerRoutes() {
	for _, route := range server.routes {
		middlewareDecoratedHandler := server.decorateHandlerWithMiddlewares(route.handler, server.middlewares)
		server.mux.HandleFunc(route.path, middlewareDecoratedHandler)
	}
}

// decorateHandlerWithMiddlewares transforms a route handler into a new function that execute all middlewares
func (server *Server) decorateHandlerWithMiddlewares(routeHandler http.HandlerFunc, middlewares []Middleware) http.HandlerFunc {
	var handler http.HandlerFunc
	for _, middleware := range middlewares {
		if handler == nil {
			handler = middleware(routeHandler)
			continue
		}
		handler = middleware(handler)
	}
	if handler == nil {
		handler = routeHandler
	}
	return handler
}

func (server *Server) acceptMethods(handler http.HandlerFunc, methods []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if compare.StringsContains(methods, r.Method) {
			handler(w, r)
		}
	}
}

// Listen starts the server in port
func (server *Server) Listen(port string) error {
	server.registerRoutes()
	return http.ListenAndServe(port, server.mux)
}
