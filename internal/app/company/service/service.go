package service

import (
	model "simple-blog-system/internal/app/company/model"
	"simple-blog-system/internal/app/company/repository"
)

type CompanyService struct {
	repo *repository.CompanyRepository
}

func NewCompanyService(repo *repository.CompanyRepository) *CompanyService {
	return &CompanyService{repo: repo}
}

func (s *CompanyService) CreateCompany(company model.Company) error {
	return s.repo.Create(company)
}

// Tambahkan di paling bawah file service.go
func (s *CompanyService) GetAllCompanies() ([]model.Company, error) {
	return s.repo.GetAll()
}

func (s *CompanyService) UpdateCompany(id string, company model.Company) error {
	return s.repo.Update(id, company)
}

func (s *CompanyService) DeleteCompany(id string) error {
	return s.repo.Delete(id)
}