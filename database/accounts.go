package database

import (
	"database/sql"
	"fmt"

	"github.com/stumbra/go-rest-bank-api/models"
)

type Storage interface {
	CreateAccount(*models.Account) error
	DeleteAccount(int) error
	UpdateAccount(*models.Account) error
	GetAccounts() ([]*models.Account, error)
	GetAccountByID(int) (*models.Account, error)
}

func (s *PostgresStore) CreateAccount(acc *models.Account) error {
	query := `INSERT INTO account
	(first_name, last_name, email, bank_id, current_balance, created_at)
	VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := s.db.Exec(
		query,
		acc.FirstName,
		acc.LastName,
		acc.Email,
		acc.BankID,
		acc.CurrentBalance,
		acc.CreatedAt)

	if err != nil {
		return err
	}

	return err
}

func (s *PostgresStore) GetAccounts() ([]*models.Account, error) {
	query := `SELECT * FROM account`

	rows, err := s.db.Query(query)

	if err != nil {
		return nil, err
	}

	accounts := []*models.Account{}

	for rows.Next() {
		account, err := scanIntoAccounts(rows)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (s *PostgresStore) GetAccount(id int) (*models.Account, error) {
	query := `SELECT * FROM account WHERE id = $1`

	rows, err := s.db.Query(query, id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccounts(rows)
	}

	return nil, fmt.Errorf("account %d not found", id)
}

func (s *PostgresStore) CreateAccountsTable() error {
	query := `CREATE TABLE IF NOT EXISTS account (
		id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
		email varchar(50),
		bank_id serial,
		current_balance serial,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func scanIntoAccounts(rows *sql.Rows) (*models.Account, error) {
	account := new(models.Account)

	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Email,
		&account.BankID,
		&account.CurrentBalance,
		&account.CreatedAt,
	)

	return account, err
}
