package handler

import (
	"net/http"

	"zond-api/internal/api/service"

	"github.com/gin-gonic/gin"
)

type ForkHandler interface {
	GetForks(c *gin.Context)
}

type forkHandler struct {
	svc service.ForkService
}

func NewForkHandler(svc service.ForkService) ForkHandler {
	return &forkHandler{svc: svc}
}

// GetForks godoc
// @Summary      Get forks
// @Description  Retrieve all recorded chain forks
// @Tags         Forks
// @Produce      json
// @Success      200  {object}  dto.ForksResponse
// @Failure      500  {object}  map[string]string "Failed to fetch forks"
// @Router       /api/forks [get]
func (h *forkHandler) GetForks(c *gin.Context) {
	forks, err := h.svc.GetForks(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch forks"})
		return
	}

	c.JSON(http.StatusOK, forks)
}
