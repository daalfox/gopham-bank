package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("gopher-bank.db"), &gorm.Config{})
	if err != nil {
		println(err)
		panic("failed to connect to database")
	}
	db.AutoMigrate(&User{})

	attachRoutes(db)

	http.ListenAndServe(":8080", nil)
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name" gorm:"unique"`
}

func attachRoutes(db *gorm.DB) {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to gopher bank!")
	})
	http.HandleFunc("GET /accounts", func(w http.ResponseWriter, r *http.Request) {
		var users []User
		db.Find(&users)
		json.NewEncoder(w).Encode(users)
	})
	http.HandleFunc("POST /register", func(w http.ResponseWriter, r *http.Request) {
		var user User

		json.NewDecoder(r.Body).Decode(&user)

		db.Create(&user)
	})
}
