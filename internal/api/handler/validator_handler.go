package handler

import (
	"net/http"

	"zond-api/internal/api/service"

	"github.com/gin-gonic/gin"
)

type ValidatorHandler interface {
	GetValidators(c *gin.Context)
}

type validatorHandler struct {
	svc service.ValidatorService
}

func NewValidatorHandler(svc service.ValidatorService) ValidatorHandler {
	return &validatorHandler{svc: svc}
}

func (h *validatorHandler) GetValidators(c *gin.Context) {
	validators, err := h.svc.GetValidators(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch validators"})
		return
	}

	c.JSON(http.StatusOK, validators)
}
