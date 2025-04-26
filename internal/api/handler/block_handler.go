package handler

import (
	"net/http"
	"strconv"

	"zond-api/internal/api/service"

	"github.com/gin-gonic/gin"
)

type BlockHandler struct {
	service service.BlockService
}

func NewBlockHandler(service service.BlockService) *BlockHandler {
	return &BlockHandler{service: service}
}

func (h *BlockHandler) GetLatestBlocks(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")
	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)
	blocks, err := h.service.GetLatestBlocks(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch blocks"})
		return
	}
	c.JSON(http.StatusOK, blocks)
}

func (h *BlockHandler) GetBlockByNumber(c *gin.Context) {
	blockNumberStr := c.Param("block_number")
	blockNumber, err := strconv.ParseInt(blockNumberStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid block number"})
		return
	}
	block, err := h.service.GetBlockByNumber(blockNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Block not found"})
		return
	}
	c.JSON(http.StatusOK, block)
}

func (h *BlockHandler) GetBlockByHash(c *gin.Context) {
	hash := c.Param("block_hash")
	block, err := h.service.GetBlockByHash(c.Request.Context(), hash)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Block not found"})
		return
	}
	c.JSON(http.StatusOK, block)
}
