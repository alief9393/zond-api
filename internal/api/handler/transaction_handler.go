package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"zond-api/internal/api/dto"
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

func (h *TransactionHandler) GetTransactionsByBlockNumber(c *gin.Context) {
	blockNumberStr := c.Param("block_number")
	blockNumber, err := strconv.ParseInt(blockNumberStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid block number"})
		return
	}

	txs, err := h.service.GetTransactionsByBlockNumber(blockNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.TransactionsResponse{Transactions: txs})
}

func (h *TransactionHandler) GetTransactionMetrics(c *gin.Context) {
	metrics, err := h.service.GetTransactionMetrics(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve transaction metrics"})
		return
	}
	c.JSON(http.StatusOK, metrics)
}

func (h *TransactionHandler) GetLatestTransactionsWithFilter(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	method := c.Query("method")
	from := c.Query("from")
	to := c.Query("to")

	txs, err := h.service.GetLatestTransactionsWithFilter(c.Request.Context(), page, limit, method, from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get filtered transactions"})
		return
	}

	var txResponses []dto.TransactionResponse
	for _, tx := range txs {
		toAddress := ""
		if len(tx.ToAddress) > 0 {
			toAddress = fmt.Sprintf("0x%x", tx.ToAddress)
		}

		txResponses = append(txResponses, dto.TransactionResponse{
			TxHash:               fmt.Sprintf("0x%x", tx.TxHash),
			BlockNumber:          tx.BlockNumber,
			FromAddress:          fmt.Sprintf("0x%x", tx.FromAddress),
			ToAddress:            toAddress,
			Value:                tx.Value,
			Gas:                  tx.Gas,
			GasPrice:             tx.GasPrice,
			Type:                 tx.Type,
			ChainID:              tx.ChainID,
			AccessList:           string(tx.AccessList),
			MaxFeePerGas:         tx.MaxFeePerGas,
			MaxPriorityFeePerGas: tx.MaxPriorityFeePerGas,
			TransactionIndex:     tx.TransactionIndex,
			CumulativeGasUsed:    tx.CumulativeGasUsed,
			IsSuccessful:         tx.IsSuccessful,
			RetrievedFrom:        tx.RetrievedFrom,
			IsCanonical:          tx.IsCanonical,
		})
	}

	c.JSON(http.StatusOK, dto.TransactionsResponse{Transactions: txResponses})
}
