package handler

import (
	"net/http"

	"zond-api/internal/api/service"

	"github.com/gin-gonic/gin"
)

type ReorgHandler interface {
	GetReorgs(c *gin.Context)
}

type reorgHandler struct {
	svc service.ReorgService
}

func NewReorgHandler(svc service.ReorgService) ReorgHandler {
	return &reorgHandler{svc: svc}
}

// GetReorgs godoc
// @Summary      Get reorgs
// @Description  Retrieve a list of chain reorganization events
// @Tags         Reorgs
// @Produce      json
// @Success      200  {object}  dto.ReorgsResponse
// @Failure      500  {object}  map[string]string "Failed to fetch reorgs"
// @Router       /api/reorgs [get]
func (h *reorgHandler) GetReorgs(c *gin.Context) {
	reorgs, err := h.svc.GetReorgs(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reorgs"})
		return
	}

	c.JSON(http.StatusOK, reorgs)
}
