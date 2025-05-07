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

// GetLatestBlocks godoc
// @Summary      Get the latest blocks
// @Description  Retrieve a paginated list of recent blocks
// @Tags         Blocks
// @Produce      json
// @Param        page    query     int  false  "Page number"     default(1)
// @Param        limit   query     int  false  "Items per page"  default(10)
// @Success      200     {object}  dto.BlocksPaginatedResponse
// @Failure      500     {object}  map[string]string "Failed to fetch blocks"
// @Router       /api/blocks/latest [get]
func (h *BlockHandler) GetLatestBlocks(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	blocks, err := h.service.GetLatestBlocks(c.Request.Context(), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch blocks"})
		return
	}
	c.JSON(http.StatusOK, blocks)
}

// GetBlockByNumber godoc
// @Summary      Get block by number
// @Description  Retrieve details of a block using its number
// @Tags         Blocks
// @Produce      json
// @Param        block_number  path      int  true  "Block number"
// @Success      200           {object}  dto.BlockResponse
// @Failure      400           {object}  map[string]string "Invalid block number"
// @Failure      404           {object}  map[string]string "Block not found"
// @Router       /api/blocks/{block_number} [get]
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

// GetBlockByHash godoc
// @Summary      Get block by hash
// @Description  Retrieve block details using its hash
// @Tags         Blocks
// @Produce      json
// @Param        hash  path      string  true  "Block hash"
// @Success      200   {object}  dto.BlockResponse
// @Failure      404   {object}  map[string]string "Block not found"
// @Router       /api/blocks/hash/{hash} [get]
func (h *BlockHandler) GetBlockByHash(c *gin.Context) {
	hash := c.Param("hash")
	block, err := h.service.GetBlockByHash(c.Request.Context(), hash)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Block not found"})
		return
	}
	c.JSON(http.StatusOK, block)
}
