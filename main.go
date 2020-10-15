package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID           int    "json:_id"
	FirstName    string "json:firstName"
	LastName     string "json:lastName"
	Instructor   bool   "json:instructor"
	Participated int    "json:participatedIn"
	Cunducted    int    "json:conducted"
}

var Users []User

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Users Endpoint Hit")
	json.NewEncoder(w).Encode(Users)
}

func main() {

	Users = []User{
		User{
			ID:           1,
			FirstName:    "Shawn",
			LastName:     "Becker",
			Instructor:   true,
			Participated: 2,
			Cunducted:    0,
		}, User{
			ID:           2,
			FirstName:    "Gordon",
			LastName:     "Becker",
			Instructor:   false,
			Participated: 0,
			Cunducted:    2,
		},
	}

	fmt.Println("Starting server on Port 8000")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// do something
	})

	http.HandleFunc("/users", getAllUsers)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
