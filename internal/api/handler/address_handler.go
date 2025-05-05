package handler

import (
	"net/http"

	"zond-api/internal/api/service"

	"github.com/gin-gonic/gin"
)

type AddressHandler interface {
	GetAddressBalance(c *gin.Context)
	GetAddressTransactions(c *gin.Context)
	GetAddressDetails(c *gin.Context)
}

type addressHandler struct {
	svc service.AddressService
}

func NewAddressHandler(svc service.AddressService) AddressHandler {
	return &addressHandler{svc: svc}
}

func (h *addressHandler) GetAddressBalance(c *gin.Context) {
	// Restrict to admin users
	role, exists := c.Get("role")
	if !exists || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
		return
	}

	address := c.Param("address")
	addr, err := h.svc.GetAddressBalance(c.Request.Context(), address)
	if err != nil {
		if err.Error() == "pgx: no rows in result set" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch balance"})
		return
	}

	c.JSON(http.StatusOK, addr)
}

func (h *addressHandler) GetAddressTransactions(c *gin.Context) {
	address := c.Param("address")
	transactions, err := h.svc.GetAddressTransactions(c.Request.Context(), address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (h *addressHandler) GetAddressDetails(c *gin.Context) {
	addr := c.Param("address")
	detail, err := h.svc.GetAddressDetails(c.Request.Context(), addr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, detail)
}
