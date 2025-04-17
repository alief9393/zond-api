package handler

import (
	"net/http"
	"strconv"

	"zond-api/internal/api/service"

	"github.com/gin-gonic/gin"
)

type BlockHandler struct {
	service *service.BlockService
}

func NewBlockHandler(service *service.BlockService) *BlockHandler {
	return &BlockHandler{service: service}
}

func (h *BlockHandler) GetLatestBlocks(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit parameter"})
		return
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil || offset < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset parameter"})
		return
	}

	response, err := h.service.GetLatestBlocks(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *BlockHandler) GetBlockByNumber(c *gin.Context) {
	blockNumberStr := c.Param("block_number")
	blockNumber, err := strconv.ParseInt(blockNumberStr, 10, 64)
	if err != nil || blockNumber < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid block number"})
		return
	}

	block, err := h.service.GetBlockByNumber(blockNumber)
	if err != nil {
		if err.Error() == "pgx: no rows in result set" {
			c.JSON(http.StatusNotFound, gin.H{"error": "block not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, block)
}
