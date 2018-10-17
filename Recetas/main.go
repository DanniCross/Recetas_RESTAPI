package main

import (
	. "RESTAPI/Recetas/Receta"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "DELETE", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router)))
}
