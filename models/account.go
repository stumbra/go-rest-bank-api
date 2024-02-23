package models

import (
	"math/rand"
	"time"
)

type Account struct {
	ID             int       `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	BankID         int64     `json:"bank_id"`
	CurrentBalance int64     `json:"current_balance"`
	CreatedAt      time.Time `json:"created_at"`
}

type CreateAccountRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func NewAccount(request CreateAccountRequest) *Account {
	return &Account{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		BankID:    int64(rand.Intn(1000000)),
		CreatedAt: time.Now().UTC(),
	}
}
