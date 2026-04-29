package service

import (
	"simple-blog-system/internal/app/bank_account/model"
	"simple-blog-system/internal/app/bank_account/repository"
)

type BankAccountService struct {
	repo *repository.BankAccountRepository
}

func NewBankAccountService(repo *repository.BankAccountRepository) *BankAccountService {
	return &BankAccountService{repo: repo}
}

func (s *BankAccountService) CreateBank(bank model.BankAccount) error {
	return s.repo.Create(bank)
}

func (s *BankAccountService) GetAllBanks() ([]model.BankAccount, error) {
	return s.repo.GetAll()
}

func (s *BankAccountService) UpdateBank(id string, bank model.BankAccount) error {
	return s.repo.Update(id, bank)
}

func (s *BankAccountService) DeleteBank(id string) error {
	return s.repo.Delete(id)
}