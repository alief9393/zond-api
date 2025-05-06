package handler

import (
	"net/http"

	"zond-api/internal/api/dto"
	"zond-api/internal/api/service"

	"github.com/gin-gonic/gin"
)

type SearchHandler struct {
	svc service.SearchService
}

func NewSearchHandler(svc service.SearchService) *SearchHandler {
	return &SearchHandler{svc: svc}
}

// GetSuggestions godoc
// @Summary      Get search suggestions
// @Description  Return suggestions based on a partial query input (e.g. address, tx hash, block number)
// @Tags         Search
// @Accept       json
// @Produce      json
// @Param        query  query     string  true  "Search input string (partial hash, address, or block number)"
// @Success      200    {object}  dto.SearchSuggestionsResponse
// @Failure      400    {object}  map[string]string "Query parameter required"
// @Failure      500    {object}  map[string]string "Internal server error"
// @Router       /api/search/suggestions [get]
func (h *SearchHandler) GetSuggestions(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter required"})
		return
	}

	suggestions, err := h.svc.GetSuggestions(c.Request.Context(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SearchSuggestionsResponse{Suggestions: suggestions})
}
