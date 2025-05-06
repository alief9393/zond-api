package handler

import (
	"net/http"
	"strconv"

	"zond-api/internal/api/dto"
	"zond-api/internal/api/service"

	"github.com/gin-gonic/gin"
)

type BeaconWithdrawalHandler struct {
	svc service.BeaconWithdrawalService
}

func NewBeaconWithdrawalHandler(svc service.BeaconWithdrawalService) *BeaconWithdrawalHandler {
	return &BeaconWithdrawalHandler{svc: svc}
}

// GetBeaconWithdrawals godoc
// @Summary      Get beacon withdrawals
// @Description  Retrieve a paginated list of beacon chain withdrawals
// @Tags         Beacon
// @Accept       json
// @Produce      json
// @Param        page   query     int  false  "Page number"        default(1)
// @Param        limit  query     int  false  "Items per page"     default(10)
// @Success      200    {object}  dto.BeaconWithdrawalsPaginatedResponse
// @Failure      500    {object}  map[string]string "Failed to retrieve beacon withdrawals"
// @Router       /beacon-withdrawals [get]
func (h *BeaconWithdrawalHandler) GetBeaconWithdrawals(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	data, total, err := h.svc.GetBeaconWithdrawals(c.Request.Context(), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve beacon withdrawals"})
		return
	}

	c.JSON(http.StatusOK, dto.BeaconWithdrawalsPaginatedResponse{
		Withdrawals: data,
		Pagination: dto.PaginationInfo{
			Page:    page,
			Limit:   limit,
			Total:   total,
			HasNext: page*limit < total,
		},
	})
}
