package handler

import (
	"net/http"

	"zond-api/internal/api/service"

	"github.com/gin-gonic/gin"
)

// AddressHandler defines address-related endpoints
type AddressHandler interface {
	GetAddressBalance(c *gin.Context)
	GetAddressTransactions(c *gin.Context)
	GetAddressDetails(c *gin.Context)
}

type addressHandler struct {
	svc service.AddressService
}

func NewAddressHandler(svc service.AddressService) AddressHandler {
	return &addressHandler{svc: svc}
}

// GetAddressBalance godoc
// @Summary      Get address balance
// @Description  Get the balance information of an address (admin only)
// @Tags         Address
// @Accept       json
// @Produce      json
// @Param        address  path      string  true  "Address"
// @Success      200      {object}  dto.AddressResponse
// @Failure      403      {object}  map[string]string "Admin access required"
// @Failure      404      {object}  map[string]string "Address not found"
// @Failure      500      {object}  map[string]string "Internal server error"
// @Security     ApiKeyAuth
// @Router       /address/{address}/balance [get]
func (h *addressHandler) GetAddressBalance(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
		return
	}

	address := c.Param("address")
	addr, err := h.svc.GetAddressBalance(c.Request.Context(), address)
	if err != nil {
		if err.Error() == "pgx: no rows in result set" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch balance"})
		return
	}

	c.JSON(http.StatusOK, addr)
}

// GetAddressTransactions godoc
// @Summary      Get address transactions
// @Description  Get list of transactions related to the given address
// @Tags         Address
// @Accept       json
// @Produce      json
// @Param        address  path      string  true  "Address"
// @Success      200      {array}   dto.TransactionDTO
// @Failure      500      {object}  map[string]string "Internal server error"
// @Router       /address/{address}/transactions [get]
func (h *addressHandler) GetAddressTransactions(c *gin.Context) {
	address := c.Param("address")
	transactions, err := h.svc.GetAddressTransactions(c.Request.Context(), address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

// GetAddressDetails godoc
// @Summary      Get address details
// @Description  Get full details of a specific address
// @Tags         Address
// @Accept       json
// @Produce      json
// @Param        address  path      string  true  "Address"
// @Success      200      {object}  dto.AddressDetailResponse
// @Failure      500      {object}  map[string]string "Internal server error"
// @Router       /address/{address}/details [get]
func (h *addressHandler) GetAddressDetails(c *gin.Context) {
	addr := c.Param("address")
	detail, err := h.svc.GetAddressDetails(c.Request.Context(), addr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, detail)
}
