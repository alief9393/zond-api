package service

import (
	"context"

	"zond-api/internal/api/dto"
	searchRepo "zond-api/internal/api/repository/search"
)

type SearchService interface {
	GetSuggestions(ctx context.Context, query string) ([]dto.Suggestion, error)
	Search(ctx context.Context, query string) ([]dto.SearchResult, error)
}
type searchService struct {
	repo searchRepo.SearchRepository
}

func NewSearchService(repo searchRepo.SearchRepository) SearchService {
	return &searchService{repo: repo}
}

func (s *searchService) GetSuggestions(ctx context.Context, query string) ([]dto.Suggestion, error) {
	return s.repo.GetSuggestions(ctx, query)
}

func (s *searchService) Search(ctx context.Context, query string) ([]dto.SearchResult, error) {
	return s.repo.Search(ctx, query)
}
