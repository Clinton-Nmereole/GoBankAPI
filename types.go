package main

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	AccountNumber int32  `json:"account_number"`
	Password      string `json:"password"`
}

type LoginResponse struct {
	Token  string `json:"token"`
	Number int32  `json:"account_number"`
}

type CreateAccountRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

type Transaction struct {
	ToAccount int32     `json:"to_account"`
	Amount    int64     `json:"amount"`
	Time      time.Time `json:"transaction_init_time"`
}

type TransactionResponse struct {
	Transaction Transaction
	Balance     int64
	Token       string
}

type Account struct {
	ID                int       `json:"id"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Number            int32     `json:"number"`
	EncryptedPassword string    `json:"-"`
	Balance           int64     `json:"balance"`
	CreatedAt         time.Time `json:"created_at"`
}

func (a *Account) PasswordMatches(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(password))
}

func NewAccount(firstName string, lastName string, password string) (*Account, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &Account{
		// ID:        rand.Intn(10000),
		FirstName:         firstName,
		LastName:          lastName,
		EncryptedPassword: string(encryptedPassword),
		Number:            rand.Int31(),
		CreatedAt:         time.Now().UTC(),
	}, nil
}
