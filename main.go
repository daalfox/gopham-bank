package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	accounts := []User{}

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to gopher bank!")
	})
	http.HandleFunc("GET /accounts", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(accounts)
	})
	http.HandleFunc("POST /register", func(w http.ResponseWriter, r *http.Request) {
		var user User
		json.NewDecoder(r.Body).Decode(&user)

		user.Id = rand.Intn(1000000)

		accounts = append(accounts, user)
	})

	http.ListenAndServe(":8080", nil)
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
