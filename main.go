package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var users []User

type User struct {
	ID    string `json:"id"`
	First string `json:"first"`
	Last  string `json:"last"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {

}

func getUser(w http.ResponseWriter, r *http.Request) {

}

func createUser(w http.ResponseWriter, r *http.Request) {

}

func updateUser(w http.ResponseWriter, r *http.Request) {

}

func deleteUser(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()

	user1 := User{
		ID:    "1",
		First: "Daniel",
		Last:  "Anatolie",
	}

	users = append(users, user1)

	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/api/users", createUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/api/users/{id}", deleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
