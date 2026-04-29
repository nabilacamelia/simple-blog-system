package model

type Company struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	LegalName string `json:"legal_name"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	LogoURL   string `json:"logo_url"`
}