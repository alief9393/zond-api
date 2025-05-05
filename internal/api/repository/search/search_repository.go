package repository

import (
	"context"
	"zond-api/internal/api/dto"
)

type SearchRepository interface {
	Search(ctx context.Context, query string) ([]dto.SearchResult, error)
	GetSuggestions(ctx context.Context, query string) ([]dto.Suggestion, error)
}
