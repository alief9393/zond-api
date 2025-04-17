package handler

import (
	"net/http"
	"strconv"

	"zond-api/internal/api/service"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	service *service.TransactionService
}

func NewTransactionHandler(service *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) GetLatestTransactions(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit parameter"})
		return
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil || offset < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset parameter"})
		return
	}

	response, err := h.service.GetLatestTransactions(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *TransactionHandler) GetTransactionByHash(c *gin.Context) {
	txHash := c.Param("tx_hash")
	if len(txHash) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "transaction hash is required"})
		return
	}

	tx, err := h.service.GetTransactionByHash(txHash)
	if err != nil {
		if err.Error() == "pgx: no rows in result set" {
			c.JSON(http.StatusNotFound, gin.H{"error": "transaction not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tx)
}
