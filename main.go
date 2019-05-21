package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/axiomzen/beanstalks-api/config"
	"github.com/axiomzen/beanstalks-api/data"
	"github.com/axiomzen/beanstalks-api/model"
	"github.com/axiomzen/beanstalks-api/server"
	"github.com/gorilla/mux"
)

//@todo: connect to SQL database
var users []User

type User struct {
	ID    string `json:"id"`
	First string `json:"first"`
	Last  string `json:"last"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, user := range users {
		if user.ID == params["id"] {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = strconv.Itoa(rand.Intn(1000000))
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, user := range users {
		if user.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			var user User
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.ID = params["id"]
			users = append(users, user)
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode(users)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, user := range users {
		if user.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}

func main() {
	fmt.Println("Starting Beanstalk API...")

	// Create server
	serv := server.New(config.FromEnv())
	serv.Start()
}

// Run tests against the DB
func test(db *data.DAL) {
	user := &model.User{
		Name:           "Bruno",
		Email:          "bruno.bachmann@dapperlabs.com",
		HashedPassword: "blablabla",
		Tags:           []string{"Back-end", "Engineering"},
	}

	db.CreateUser(user)
}
