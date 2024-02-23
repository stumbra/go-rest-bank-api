package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

func (s *PostgresStore) Init() error {
	return s.CreateAccountsTable()
}

func NewPostgresStore() *PostgresStore {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", os.Getenv("DB_USERNAME"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	if err == nil {
		fmt.Println("Succesfully connected with DB...")
	}

	store := &PostgresStore{
		db: db,
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	return store
}
