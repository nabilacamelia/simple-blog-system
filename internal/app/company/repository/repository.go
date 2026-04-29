package repository

import (
	"database/sql"
	model "simple-blog-system/internal/app/company/model"
)

type CompanyRepository struct {
	db *sql.DB
}

func NewCompanyRepository(db *sql.DB) *CompanyRepository {
	return &CompanyRepository{db: db}
}

func (r *CompanyRepository) Create(company model.Company) error {
	query := `INSERT INTO companies (name, legal_name, address, phone, email, logo_url) 
              VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(query, company.Name, company.LegalName, company.Address, company.Phone, company.Email, company.LogoURL)
	return err
}

// Tambahkan di paling bawah file repository.go
func (r *CompanyRepository) GetAll() ([]model.Company, error) {
	query := `SELECT id, name, legal_name, address, phone, email, logo_url FROM companies`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companies []model.Company
	for rows.Next() {
		var c model.Company
		err := rows.Scan(&c.ID, &c.Name, &c.LegalName, &c.Address, &c.Phone, &c.Email, &c.LogoURL)
		if err != nil {
			return nil, err
		}
		companies = append(companies, c)
	}
	return companies, nil
}

func (r *CompanyRepository) Update(id string, company model.Company) error {
	query := `UPDATE companies 
              SET name=$1, legal_name=$2, address=$3, phone=$4, email=$5, logo_url=$6 
              WHERE id=$7`
	_, err := r.db.Exec(query, company.Name, company.LegalName, company.Address, company.Phone, company.Email, company.LogoURL, id)
	return err
}

func (r *CompanyRepository) Delete(id string) error {
	query := `DELETE FROM companies WHERE id=$1`
	_, err := r.db.Exec(query, id)
	return err
}