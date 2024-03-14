package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Users struct {
	Name string `json:"name"`
}

var UserList []string

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(UserList)
}

func main() {
	http.HandleFunc("/user", getUsers)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
