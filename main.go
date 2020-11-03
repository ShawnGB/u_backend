package main

import (
	"log"
	"net/http"
	"u_admin/backend"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	backend.InitDb()

	router := backend.NewRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}

// db functions
