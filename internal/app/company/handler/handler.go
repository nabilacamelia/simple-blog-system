package handler

import (
	"net/http"
	model "simple-blog-system/internal/app/company/model"
	"simple-blog-system/internal/app/company/service"
	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	service *service.CompanyService
}

func NewCompanyHandler(service *service.CompanyService) *CompanyHandler {
	return &CompanyHandler{service: service}
}

func (h *CompanyHandler) CreateCompany(c *gin.Context) {
	var payload model.Company
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateCompany(payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal simpan data"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Data perusahaan berhasil ditambahkan", "data": payload})
}

// Tambahkan di paling bawah file handler.go
func (h *CompanyHandler) GetCompanies(c *gin.Context) {
	companies, err := h.service.GetAllCompanies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil mengambil data perusahaan",
		"data":    companies,
	})
}

func (h *CompanyHandler) UpdateCompany(c *gin.Context) {
	id := c.Param("id") // Mengambil ID dari URL
	var payload model.Company
	
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateCompany(id, payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data perusahaan berhasil diperbarui"})
}

func (h *CompanyHandler) DeleteCompany(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteCompany(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data perusahaan berhasil dihapus"})
}