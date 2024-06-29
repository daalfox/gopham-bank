package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatal("database connection attempt failed")
	}

	bank := newBank(db)
	bank.run()
}

func newBank(db *gorm.DB) GophamBank {
	return GophamBank{
		db,
	}
}

type GophamBank struct {
	db *gorm.DB
}

func (b *GophamBank) run() {
	// run the service and listen for events
}
func (b *GophamBank) register(username, password string) {
	newAccount := Account{
		Username: username,
		Password: password,
	}
	b.db.Create(&newAccount)
	fmt.Println("new account created:")
	fmt.Printf("ID: %v\n", newAccount.ID)
	fmt.Printf("username: %v\n", newAccount.Username)
	fmt.Printf("balance: %0.2f\n", newAccount.Balance)
}
func (b *GophamBank) login(username, password string) {
	var result Account

	b.db.Where("username = ? AND password = ?", username, password).First(&result)
	fmt.Printf("`%v (ID %v)` has `%0.2f` in his/her account\n", result.Username, result.ID, result.Balance)
}

type Account struct {
	ID                 uint
	Username, Password string
	Balance            float64
}

func (b *GophamBank) withdraw() {
	// Somehow withdraw funds
}
func (b *GophamBank) deposit() {
	// Somehow deposit funds
}
