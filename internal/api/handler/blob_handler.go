package handler

import (
	"net/http"
	"strconv"

	"zond-api/internal/api/dto"
	"zond-api/internal/api/service"

	"github.com/gin-gonic/gin"
)

type BlobHandler struct {
	svc service.BlobService
}

func NewBlobHandler(svc service.BlobService) *BlobHandler {
	return &BlobHandler{svc: svc}
}

// GetBlobs godoc
// @Summary      Get blobs
// @Description  Retrieve a paginated list of blobs
// @Tags         Blobs
// @Accept       json
// @Produce      json
// @Param        page   query     int  false  "Page number"  default(1)
// @Param        limit  query     int  false  "Items per page"  default(10)
// @Success      200    {object}  dto.BlobsPaginatedResponse
// @Failure      500    {object}  map[string]string
// @Router       /api/blobs [get]
func (h *BlobHandler) GetBlobs(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	blobs, total, err := h.svc.GetBlobs(c.Request.Context(), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch blobs"})
		return
	}

	c.JSON(http.StatusOK, dto.BlobsPaginatedResponse{
		Blobs: blobs,
		Pagination: dto.PaginationInfo{
			Page:  page,
			Limit: limit,
			Total: total,
		},
	})
}
