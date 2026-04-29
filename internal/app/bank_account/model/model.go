package model

type BankAccount struct {
	ID            int    `json:"id"`
	BankName      string `json:"bank_name"`
	AccountNumber string `json:"account_number"`
	AccountHolder string `json:"account_holder"`
}