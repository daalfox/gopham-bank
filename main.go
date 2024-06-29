package main

import (
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
func (b *GophamBank) createAccount(username, password string) {
	// Somehow create account for user
}
func (b *GophamBank) withdraw() {
	// Somehow withdraw funds
}
func (b *GophamBank) deposit() {
	// Somehow deposit funds
}
