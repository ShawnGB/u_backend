package main

import (
	"log"
	"net/http"
	"u_admin/backend"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
)

func main() {

	backend.Initdb()

	router := backend.NewRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})

	log.Fatal(http.ListenAndServe("localhost:8000",
		handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)))
}

// db functions
