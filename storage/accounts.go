package storage

import "github.com/stumbra/go-rest-bank-api/models"

type Storage interface {
	CreateAccount(*models.Account) error
	DeleteAccount(int) error
	UpdateAccount(*models.Account) error
	GetAccountByID(int) (*models.Account, error)
}
