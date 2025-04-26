package service

import (
	"context"

	"zond-api/internal/api/dto"
	"zond-api/internal/api/repository"
)

type SearchService struct {
	repo repository.SearchRepository
}

func NewSearchService(repo repository.SearchRepository) *SearchService {
	return &SearchService{repo: repo}
}

func (s *SearchService) Search(ctx context.Context, query string) ([]dto.SearchResult, error) {
	return s.repo.Search(ctx, query)
}
