package repository

import (
	"database/sql"
	"simple-blog-system/internal/app/bank_account/model"
)

type BankAccountRepository struct {
	db *sql.DB
}

func NewBankAccountRepository(db *sql.DB) *BankAccountRepository {
	return &BankAccountRepository{db: db}
}

func (r *BankAccountRepository) Create(bank model.BankAccount) error {
	query := `INSERT INTO bank_accounts (bank_name, account_number, account_holder) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, bank.BankName, bank.AccountNumber, bank.AccountHolder)
	return err
}

func (r *BankAccountRepository) GetAll() ([]model.BankAccount, error) {
	rows, err := r.db.Query("SELECT id, bank_name, account_number, account_holder FROM bank_accounts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []model.BankAccount
	for rows.Next() {
		var a model.BankAccount
		rows.Scan(&a.ID, &a.BankName, &a.AccountNumber, &a.AccountHolder)
		accounts = append(accounts, a)
	}
	return accounts, nil
}

func (r *BankAccountRepository) Update(id string, bank model.BankAccount) error {
	query := `UPDATE bank_accounts SET bank_name=$1, account_number=$2, account_holder=$3 WHERE id=$4`
	_, err := r.db.Exec(query, bank.BankName, bank.AccountNumber, bank.AccountHolder, id)
	return err
}

func (r *BankAccountRepository) Delete(id string) error {
	query := `DELETE FROM bank_accounts WHERE id=$1`
	_, err := r.db.Exec(query, id)
	return err
}