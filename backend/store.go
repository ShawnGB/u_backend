package backend

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// connect and open db
func Initdb() *sql.DB {
	var err error
	db, err = sql.Open("mysql", "root@/uskillity")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	result, err := db.Query("select * from users")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var user User
		err := result.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Instructor)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("INSERT INTO users(FirstName, LastName, Instructor) VALUES(?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	var user User
	json.Unmarshal(body, &user)
	_, err = stmt.Exec(user.Firstname, user.Lastname, boolToInt(user.Instructor))
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New user was created")
}

func boolToInt(x bool) int {
	if x {
		return 1
	} else {
		return 0
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT UserId, FirstName, LastName FROM users WHERE UserId = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var user User
	for result.Next() {
		err := result.Scan(&user.ID, &user.Firstname, &user.Lastname)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE users SET FirstName = ? WHERE _id = ?")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	newName := keyVal["firstName"]
	_, err = stmt.Exec(newName, params["_id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "User with ID = %s was updated", params["_id"])
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM users WHERE _id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["_id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "User with ID = %s was deleted", params["_id"])
}
