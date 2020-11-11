package main

import (
	"log"
	"net/http"
	"u_admin/backend"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	backend.Initdb()

	router := backend.NewRouter()

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

// db functions
