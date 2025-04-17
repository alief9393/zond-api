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

func (h *reorgHandler) GetReorgs(c *gin.Context) {
	reorgs, err := h.svc.GetReorgs(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reorgs"})
		return
	}

	c.JSON(http.StatusOK, reorgs)
}
