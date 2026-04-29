package server

import (
	"github.com/gin-gonic/gin"

	"simple-blog-system/internal/app/user/port"
	compHandler "simple-blog-system/internal/app/company/handler"
	bankHandler "simple-blog-system/internal/app/bank_account/handler" // Pastikan import ini ada
)

type routes struct{}

var Routes routes

func (r routes) New(router *gin.RouterGroup, handler port.IUserHandler, companyHandler *compHandler.CompanyHandler, bankAccountHandler *bankHandler.BankAccountHandler) {
	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)

	// --- Jalur CRUD Company ---
	router.POST("/company", companyHandler.CreateCompany)
	router.GET("/company", companyHandler.GetCompanies)
	router.PUT("/company/:id", companyHandler.UpdateCompany)
	router.DELETE("/company/:id", companyHandler.DeleteCompany)

	// --- Jalur CRUD Bank Account (PASTIKAN DI DALAM KURUNG INI) ---
	router.POST("/bank-account", bankAccountHandler.CreateBank)
	router.GET("/bank-account", bankAccountHandler.GetBanks)
	router.PUT("/bank-account/:id", bankAccountHandler.UpdateBank)
	router.DELETE("/bank-account/:id", bankAccountHandler.DeleteBank)
} // <--- Penutup fungsi New

func (r routes) NewProfile(router *gin.RouterGroup, handler port.IUserHandler) {
	router.GET("/", handler.GetUser)
	router.PUT("", handler.UpdateUser)
}