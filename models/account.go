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

func NewAccount(firstName, lastName, email string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		BankID:    int64(rand.Intn(1000000)),
		CreatedAt: time.Now().UTC(),
	}
}

type CreateAccountRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
}
