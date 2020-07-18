package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

//Related to error "should not use basic type string as key in context.WithValue"
type contextKey string

var userContextKey contextKey = "user"
var requestIDKey contextKey = "requestID"

func main() {
	server := NewServer()
	server.UseMiddleware(traceRequest)

	server.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/", r.Context().Value(requestIDKey))
		fmt.Fprintf(w, "essa eh do server default")
	})

	router := NewRouter()
	router.GET("/a", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/a", r.Context().Value(requestIDKey))
		fmt.Fprintf(w, "essa eh do router")
	})

	server.UseRouter(router)

	log.Fatal(server.Listen(":8080"))
}

func traceRequest(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		context := context.WithValue(r.Context(), requestIDKey, time.Now().String())
		f(w, r.WithContext(context))
	}
}
