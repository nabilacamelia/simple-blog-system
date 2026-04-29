package handler

import (
	"net/http"
	"simple-blog-system/internal/app/bank_account/model"
	"simple-blog-system/internal/app/bank_account/service"
	"github.com/gin-gonic/gin"
)

type BankAccountHandler struct {
	service *service.BankAccountService
}

func NewBankAccountHandler(s *service.BankAccountService) *BankAccountHandler {
	return &BankAccountHandler{service: s}
}

func (h *BankAccountHandler) CreateBank(c *gin.Context) {
	var payload model.BankAccount
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.service.CreateBank(payload)
	c.JSON(http.StatusOK, gin.H{"message": "Bank Account berhasil ditambahkan"})
}

func (h *BankAccountHandler) GetBanks(c *gin.Context) {
	banks, _ := h.service.GetAllBanks()
	c.JSON(http.StatusOK, gin.H{"data": banks})
}

func (h *BankAccountHandler) UpdateBank(c *gin.Context) {
	id := c.Param("id")
	var payload model.BankAccount
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.UpdateBank(id, payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update bank"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Bank Account berhasil diperbarui"})
}

func (h *BankAccountHandler) DeleteBank(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteBank(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal hapus bank"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Bank Account berhasil dihapus"})
}