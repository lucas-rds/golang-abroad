package main

import (
	"abroad/application/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := NewServer()
	server.AddMiddleware(func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Middleware Executando", r.URL)
			f(w, r)
		}
	})

	server.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Get principal")
		ag := models.AgencyResponse{
			Name: "José",
		}
		agJSON, _ := json.Marshal(ag)
		fmt.Fprintf(w, "Hello, %s", agJSON)
	})

	server.GET("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {})

	log.Fatal(server.Listen(":8080"))
}
