package server

import (
	"fmt"
	"net/http"
)

/*
	Trying to encapsulate and add functionality into default golang server
	Also trying to keep it simple.
*/

// Middleware is a function that mas execute next to continue the request flow
type Middleware func(next http.HandlerFunc) http.HandlerFunc

// Server is an net/http server with
type Server struct {
	mux         *http.ServeMux
	router      *Router
	middlewares []Middleware
}

// NewServer is a Server builder
func NewServer() *Server {
	return &Server{
		mux:         http.NewServeMux(),
		router:      NewRouter(),
		middlewares: make([]Middleware, 0),
	}
}

// GET adds a get HandleFunc into server mux
func (server *Server) GET(path string, handler http.HandlerFunc) {
	server.router.GET(path, handler)
}

// POST adds a post HandleFunc into server mux
func (server *Server) POST(path string, handler http.HandlerFunc) {
	server.router.POST(path, handler)
}

// AddRoute adds a Route into server mux
func (server *Server) AddRoute(path string, handler http.HandlerFunc, methods []string) {
	server.router.AddRoute(path, handler, methods)
}

// UseMiddleware adds a Middleware into server to decorate all routes
func (server *Server) UseMiddleware(callback Middleware) {
	server.middlewares = append(server.middlewares, callback)
}

// UseRouter adds all router routes into main server
func (server *Server) UseRouter(router *Router) {
	for _, route := range router.routes {
		for _, alreadyRegisteredRoute := range server.router.routes {
			if route.path == alreadyRegisteredRoute.path {
				panic(fmt.Sprintf("Route path already registered %s", route.path))
			}
		}
		server.router.routes = append(server.router.routes, route)
	}
}

// decorateHandlerWithMiddlewares transforms a route handler into a new function that execute all middlewares keeping the AddMiddleware order
func (server *Server) decorateHandlerWithMiddlewares(routeHandler http.HandlerFunc, middlewares []Middleware) http.HandlerFunc {
	handler := routeHandler
	for index := len(middlewares) - 1; index >= 0; index-- {
		handler = middlewares[index](handler)
	}
	return handler
}

func (server *Server) registerRoutes() {
	for _, route := range server.router.routes {
		middlewareDecoratedHandler := server.decorateHandlerWithMiddlewares(route.handler, server.middlewares)
		server.mux.HandleFunc(route.path, middlewareDecoratedHandler)
	}
}

// Listen starts the server in port
func (server *Server) Listen(port string) error {
	server.registerRoutes()
	return http.ListenAndServe(port, server.mux)
}
