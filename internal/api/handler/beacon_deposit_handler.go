package handler

import (
	"net/http"
	"strconv"

	"zond-api/internal/api/dto"
	"zond-api/internal/api/service"

	"github.com/gin-gonic/gin"
)

type BeaconDepositHandler struct {
	svc service.BeaconDepositService
}

func NewBeaconDepositHandler(svc service.BeaconDepositService) *BeaconDepositHandler {
	return &BeaconDepositHandler{svc: svc}
}

// GetBeaconDeposits godoc
// @Summary      Get beacon deposits
// @Description  Retrieve a paginated list of beacon chain deposits
// @Tags         Beacon
// @Accept       json
// @Produce      json
// @Param        page   query     int  false  "Page number"  default(1)
// @Param        limit  query     int  false  "Items per page"  default(10)
// @Success      200    {object}  dto.BeaconDepositsPaginatedResponse
// @Failure      500    {object}  map[string]string "Failed to fetch beacon deposits"
// @Router       /beacon-deposits [get]
func (h *BeaconDepositHandler) GetBeaconDeposits(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	deposits, total, err := h.svc.GetBeaconDeposits(c.Request.Context(), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch beacon deposits"})
		return
	}

	c.JSON(http.StatusOK, dto.BeaconDepositsPaginatedResponse{
		Deposits: deposits,
		Pagination: dto.PaginationInfo{
			Page:    page,
			Limit:   limit,
			Total:   total,
			HasNext: page*limit < total,
		},
	})
}
