package handler

import (
	"net/http"
	"strconv"

	"zond-api/internal/api/service"

	"github.com/gin-gonic/gin"
)

type ValidatorHandler interface {
	GetValidators(c *gin.Context)
	GetValidatorByID(c *gin.Context)
}

type validatorHandler struct {
	svc service.ValidatorService
}

func NewValidatorHandler(svc service.ValidatorService) ValidatorHandler {
	return &validatorHandler{svc: svc}
}

// GetValidators godoc
// @Summary      Get validators
// @Description  Retrieve a list of all active validators
// @Tags         Validators
// @Produce      json
// @Success      200  {object}  dto.ValidatorsResponse
// @Failure      500  {object}  map[string]string "Failed to fetch validators"
// @Router       /api/validators [get]
func (h *validatorHandler) GetValidators(c *gin.Context) {
	validators, err := h.svc.GetValidators(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch validators"})
		return
	}

	c.JSON(http.StatusOK, validators)
}

// GetValidatorByID godoc
// @Summary      Get validator by ID
// @Description  Retrieve a specific validator by its index
// @Tags         Validators
// @Produce      json
// @Param        validatorId  path      int  true  "Validator index"
// @Success      200          {object}  dto.ValidatorDetailResponse
// @Failure      400          {object}  map[string]string "Invalid validator index"
// @Failure      404          {object}  map[string]string "Validator not found"
// @Router       /api/validators/{validatorId} [get]
func (h *validatorHandler) GetValidatorByID(c *gin.Context) {
	indexStr := c.Param("validatorId")
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid validator index"})
		return
	}

	validator, err := h.svc.GetValidatorByID(c.Request.Context(), index)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Validator not found"})
		return
	}

	c.JSON(http.StatusOK, validator)
}
