package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(account *Account) error
	GetAccountByID(id int) (*Account, error)
	GetAccountByNumber(number int32) (*Account, error)
	DeleteAccount(id int) error
	GetAllAccounts() ([]*Account, error)
	UpdateAccount(account *Account) error
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() (*PostgresStorage, error) {
	connectString := "user=postgres password=Abrightdayin1990 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connectString)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStorage{
		db: db,
	}, nil
}

func (s *PostgresStorage) Init() error {
	return s.CreateAccountTable()
}

func (s *PostgresStorage) CreateAccountTable() error {
	query := "CREATE TABLE IF NOT EXISTS accounts (id SERIAL PRIMARY KEY, first_name VARCHAR(255), last_name VARCHAR(255), number BIGINT, balance BIGINT, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP)"
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStorage) CreateAccount(account *Account) error {
	_, err := s.db.Exec(
		"INSERT INTO accounts (first_name, last_name, number, balance, created_at) VALUES ($1, $2, $3, $4, $5)",
		account.FirstName,
		account.LastName,
		account.Number,
		account.Balance,
		account.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStorage) GetAccountByID(id int) (*Account, error) {
	row, _ := s.db.Query("SELECT * FROM accounts WHERE id = $1", id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	for row.Next() {
		return scanIntoAccount(row)
	}
	return nil, fmt.Errorf("account %d not found", id)
}

func (s *PostgresStorage) GetAccountByNumber(number int32) (*Account, error) {
	row, _ := s.db.Query("SELECT * FROM accounts WHERE number = $1", number)
	if row.Err() != nil {
		return nil, row.Err()
	}

	for row.Next() {
		return scanIntoAccount(row)
	}
	return nil, fmt.Errorf("account with number [%d] not found", number)
}

func (s *PostgresStorage) DeleteAccount(id int) error {
	_, err := s.db.Exec("DELETE FROM accounts WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStorage) UpdateAccount(account *Account) error {
	return nil
}

func (s *PostgresStorage) GetAllAccounts() ([]*Account, error) {
	rows, _ := s.db.Query("SELECT * FROM accounts")
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	var accounts []*Account
	for rows.Next() {
		var account Account
		if err := rows.Scan(&account.ID, &account.FirstName, &account.LastName, &account.Number, &account.Balance, &account.CreatedAt); err != nil {
			return nil, err
		}
		accounts = append(accounts, &account)
	}
	return accounts, nil
}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	if err := rows.Scan(&account.ID, &account.FirstName, &account.LastName, &account.Number, &account.Balance, &account.CreatedAt); err != nil {
		return nil, err
	}
	return account, nil
}
