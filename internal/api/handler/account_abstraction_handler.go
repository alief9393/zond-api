package handler

import (
	"net/http"
	"strconv"

	"zond-api/internal/api/dto"
	"zond-api/internal/api/service"

	"github.com/gin-gonic/gin"
)

type AccountAbstractionHandler struct {
	svc service.AccountAbstractionService
}

func NewAccountAbstractionHandler(svc service.AccountAbstractionService) *AccountAbstractionHandler {
	return &AccountAbstractionHandler{svc: svc}
}

// GetAccountAbstraction godoc
// @Summary      Get AA and Bundle Transactions
// @Description  Retrieve paginated Account Abstraction data (AA Txns & Bundles)
// @Tags         Account Abstraction
// @Accept       json
// @Produce      json
// @Param        page   query     int  false  "Page number"  default(1)
// @Param        limit  query     int  false  "Items per page"  default(10)
// @Success      200    {object}  dto.AccountAbstractionPaginatedResponse
// @Failure      500    {object}  map[string]string
// @Router       /api/account-abstraction [get]
func (h *AccountAbstractionHandler) GetAccountAbstraction(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	data, total, err := h.svc.GetAccountAbstraction(c.Request.Context(), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch account abstraction data"})
		return
	}

	c.JSON(http.StatusOK, dto.AccountAbstractionPaginatedResponse{
		Data: data,
		Pagination: dto.PaginationInfo{
			Page:  page,
			Limit: limit,
			Total: total,
		},
	})
}

// GetOnlyBundleTransactions godoc
// @Summary      Get Bundle Transactions only
// @Description  Retrieve paginated bundle transactions
// @Tags         Account Abstraction
// @Accept       json
// @Produce      json
// @Param        page   query     int  false  "Page number"  default(1)
// @Param        limit  query     int  false  "Items per page"  default(10)
// @Success      200    {object}  dto.BundlesPaginatedResponse
// @Failure      500    {object}  map[string]string
// @Router       /api/account-abstraction/bundles [get]
func (h *AccountAbstractionHandler) GetOnlyBundleTransactions(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	data, total, err := h.svc.GetBundlesOnly(c.Request.Context(), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch bundle transactions"})
		return
	}

	c.JSON(http.StatusOK, dto.BundlesPaginatedResponse{
		Bundles: data,
		Pagination: dto.PaginationInfo{
			Page:  page,
			Limit: limit,
			Total: total,
		},
	})
}

// GetOnlyAATransactions godoc
// @Summary      Get AA Transactions only
// @Description  Retrieve paginated AA transactions
// @Tags         Account Abstraction
// @Accept       json
// @Produce      json
// @Param        page   query     int  false  "Page number"  default(1)
// @Param        limit  query     int  false  "Items per page"  default(10)
// @Success      200    {object}  dto.AATransactionsPaginatedResponse
// @Failure      500    {object}  map[string]string
// @Router       /api/account-abstraction/aa [get]
func (h *AccountAbstractionHandler) GetOnlyAATransactions(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	data, total, err := h.svc.GetAATransactionsOnly(c.Request.Context(), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch AA transactions"})
		return
	}

	c.JSON(http.StatusOK, dto.AATransactionsPaginatedResponse{
		AATransactions: data,
		Pagination: dto.PaginationInfo{
			Page:  page,
			Limit: limit,
			Total: total,
		},
	})
}
