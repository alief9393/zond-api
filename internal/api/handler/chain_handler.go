package handler

import (
	"net/http"

	"zond-api/internal/api/service"

	"github.com/gin-gonic/gin"
)

type ChainHandler interface {
	GetChainInfo(c *gin.Context)
}

type chainHandler struct {
	svc service.ChainService
}

func NewChainHandler(svc service.ChainService) ChainHandler {
	return &chainHandler{svc: svc}
}

// GetChainInfo godoc
// @Summary      Get chain info
// @Description  Retrieve general information about the current chain
// @Tags         Chain
// @Produce      json
// @Success      200  {object}  dto.ChainResponse
// @Failure      500  {object}  map[string]string "Failed to fetch chain info"
// @Router       /api/chain [get]
func (h *chainHandler) GetChainInfo(c *gin.Context) {
	chain, err := h.svc.GetChainInfo(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch chain info"})
		return
	}

	c.JSON(http.StatusOK, chain)
}
