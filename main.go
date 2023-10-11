package main

import (
	"flag"
	"fmt"
	"log"
)

func seedAccount(store Storage, firstname, lastname, password string) *Account {
	account, err := NewAccount(firstname, lastname, password)
	if err != nil {
		log.Fatal(err)
	}
	if err := store.CreateAccount(account); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Account created => ", account.Number)

	return account
}

func seedAccounts(store Storage) {
	seedAccount(store, "Code", "Sensei", "password")
}

func main() {
	seed := flag.Bool("seed", false, "Seed the database")
	flag.Parse()

	store, err := NewPostgresStorage()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Println("Seeding database")
		seedAccounts(store)
	}

	server := NewAPIServer("localhost:8080", store)
	server.Start()
}
