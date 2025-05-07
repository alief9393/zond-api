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

// GetLatestTransactions godoc
// @Summary      Get latest transactions
// @Description  Retrieve the latest transactions with pagination
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Param        limit   query     int  false  "Number of items per page"
// @Param        offset  query     int  false  "Offset from the beginning"
// @Success      200     {object}  dto.TransactionsResponse
// @Failure      400     {object}  ErrorResponse
// @Failure      500     {object}  ErrorResponse
// @Router       /transactions/latest [get]
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

// GetTransactionByHash godoc
// @Summary      Get transaction by hash
// @Description  Retrieve transaction details using the transaction hash
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Param        tx_hash  path      string  true  "Transaction Hash"
// @Success      200      {object}  dto.TransactionResponse
// @Failure      400      {object}  ErrorResponse
// @Failure      404      {object}  ErrorResponse
// @Router       /transactions/{tx_hash} [get]
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

// GetTransactionsByBlockNumber godoc
// @Summary      Get transactions by block number
// @Description  Retrieve transactions included in a specific block
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Param        block_number  path      int  true  "Block Number"
// @Success      200           {object}  dto.TransactionsResponse
// @Failure      400           {object}  ErrorResponse
// @Failure      500           {object}  ErrorResponse
// @Router       /transactions/block/{block_number} [get]
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

// GetTransactionMetrics godoc
// @Summary      Get transaction metrics
// @Description  Fetch aggregated metrics about transactions
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.TransactionMetricsResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /transactions/metrics [get]
func (h *TransactionHandler) GetTransactionMetrics(c *gin.Context) {
	metrics, err := h.service.GetTransactionMetrics(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve transaction metrics"})
		return
	}
	c.JSON(http.StatusOK, metrics)
}

// GetLatestTransactionsWithFilter godoc
// @Summary      Get filtered transactions
// @Description  Retrieve latest transactions with optional method/from/to filters
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Param        page    query     int     false  "Page number"
// @Param        limit   query     int     false  "Limit per page"
// @Param        method  query     string  false  "Method name"
// @Param        from    query     string  false  "From address"
// @Param        to      query     string  false  "To address"
// @Success      200     {object}  dto.TransactionsPaginatedResponse
// @Failure      500     {object}  ErrorResponse
// @Router       /transactions [get]
func (h *TransactionHandler) GetLatestTransactionsWithFilter(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	method := c.Query("method")
	from := c.Query("from")
	to := c.Query("to")

	ctx := c.Request.Context()

	txs, err := h.service.GetLatestTransactionsWithFilter(ctx, page, limit, method, from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get filtered transactions"})
		return
	}

	total, err := h.service.CountTransactionsWithFilter(ctx, method, from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to count transactions"})
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

	c.JSON(http.StatusOK, dto.TransactionsPaginatedResponse{
		Transactions: txResponses,
		Pagination: dto.PaginationInfo{
			Page:  page,
			Limit: limit,
			Total: total,
		},
	})
}

// GetPendingTransactions godoc
// @Summary      Get pending transactions
// @Description  Retrieve list of pending transactions
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Param        page    query     int     false  "Page number"
// @Param        limit   query     int     false  "Limit per page"
// @Param        method  query     string  false  "Method name"
// @Param        from    query     string  false  "From address"
// @Param        to      query     string  false  "To address"
// @Success      200     {object}  dto.TransactionsPaginatedResponse
// @Failure      500     {object}  ErrorResponse
// @Router       /transactions/pending [get]
func (h *TransactionHandler) GetPendingTransactions(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	method := c.Query("method")
	from := c.Query("from")
	to := c.Query("to")

	transactions, count, err := h.service.GetPendingTransactions(c.Request.Context(), page, limit, method, from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve pending transactions"})
		return
	}

	var txResponses []dto.TransactionResponse
	for _, tx := range transactions {
		txResponses = append(txResponses, dto.TransactionResponse{
			TxHash:               fmt.Sprintf("0x%x", tx.TxHash),
			BlockNumber:          tx.BlockNumber,
			FromAddress:          fmt.Sprintf("0x%x", tx.FromAddress),
			ToAddress:            fmt.Sprintf("0x%x", tx.ToAddress),
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

	c.JSON(http.StatusOK, dto.TransactionsPaginatedResponse{
		Transactions: txResponses,
		Pagination: dto.PaginationInfo{
			Page:  page,
			Limit: limit,
			Total: count,
		},
	})

}

// GetContractTransactions godoc
// @Summary      Get contract transactions
// @Description  Retrieve contract interactions with filters and pagination
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Param        page    query     int     false  "Page number"
// @Param        limit   query     int     false  "Limit per page"
// @Param        method  query     string  false  "Method name"
// @Param        from    query     string  false  "From address"
// @Param        to      query     string  false  "To address"
// @Success      200     {object}  dto.TransactionsPaginatedResponse
// @Failure      500     {object}  ErrorResponse
// @Router       /transactions/contract [get]
func (h *TransactionHandler) GetContractTransactions(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	method := c.Query("method")
	from := c.Query("from")
	to := c.Query("to")

	ctx := c.Request.Context()

	txs, err := h.service.GetContractTransactions(ctx, page, limit, method, from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get contract transactions"})
		return
	}

	count, err := h.service.CountContractTransactions(ctx, method, from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to count contract transactions"})
		return
	}

	var txResponses []dto.TransactionResponse
	for _, tx := range txs {
		txResponses = append(txResponses, dto.TransactionResponse{
			TxHash:               fmt.Sprintf("0x%x", tx.TxHash),
			BlockNumber:          tx.BlockNumber,
			FromAddress:          fmt.Sprintf("0x%x", tx.FromAddress),
			ToAddress:            fmt.Sprintf("0x%x", tx.ToAddress),
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

	c.JSON(http.StatusOK, dto.TransactionsPaginatedResponse{
		Transactions: txResponses,
		Pagination: dto.PaginationInfo{
			Page:  page,
			Limit: limit,
			Total: count,
		},
	})
}

// GetDailyTransactionStats godoc
// @Summary      Get daily transaction stats
// @Description  Returns the number of transactions per day for the last N days
// @Tags         Transactions
// @Produce      json
// @Param        days  query     int  false  "Number of days" default(14)
// @Success      200   {object}  dto.DailyTransactionStatsResponse
// @Failure      400   {object}  dto.ErrorResponse
// @Failure      500   {object}  dto.ErrorResponse
// @Router       /transactions/stats/daily [get]
func (h *TransactionHandler) GetDailyTransactionStats(c *gin.Context) {
	daysStr := c.DefaultQuery("days", "14")
	days, err := strconv.Atoi(daysStr)
	if err != nil || days <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid days parameter"})
		return
	}

	stats, err := h.service.GetDailyTransactionStats(days)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch daily stats"})
		return
	}

	c.JSON(http.StatusOK, dto.DailyTransactionStatsResponse{Data: stats})
}

// GetTPSStats godoc
// @Summary      Get average TPS
// @Description  Calculate average transactions per second based on latest blocks
// @Tags         Transactions
// @Produce      json
// @Param        blocks  query     int  false  "Number of recent blocks to calculate TPS from" default(100)
// @Success      200     {object}  dto.TPSStatResponse
// @Failure      500     {object}  dto.ErrorResponse
// @Router       /transactions/stats/tps [get]
func (h *TransactionHandler) GetTPSStats(c *gin.Context) {
	blockCountStr := c.DefaultQuery("blocks", "100")
	blockCount, err := strconv.Atoi(blockCountStr)
	if err != nil || blockCount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid blocks parameter"})
		return
	}

	tps, err := h.service.GetAverageTPS(blockCount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to calculate TPS"})
		return
	}

	c.JSON(http.StatusOK, tps)
}

// GetDailyFeeStats godoc
// @Summary      Get daily fee stats
// @Description  Returns total fee in ETH and average fee in USD per day
// @Tags         Transactions
// @Produce      json
// @Param        days  query     int  false  "Number of days to include" default(14)
// @Success      200   {object}  dto.DailyFeeStatsResponse
// @Failure      400   {object}  dto.ErrorResponse
// @Failure      500   {object}  dto.ErrorResponse
// @Router       /transactions/stats/fee/daily [get]
func (h *TransactionHandler) GetDailyFeeStats(c *gin.Context) {
	daysStr := c.DefaultQuery("days", "14")
	days, err := strconv.Atoi(daysStr)
	if err != nil || days <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid days parameter"})
		return
	}

	stats, err := h.service.GetDailyFeeStats(days)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve fee stats"})
		return
	}

	c.JSON(http.StatusOK, dto.DailyFeeStatsResponse{Data: stats})
}
