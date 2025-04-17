package service

import (
	"context"
	"zond-api/internal/api/dto"
	"zond-api/internal/api/repository"
)

type ReorgService interface {
	GetReorgs(ctx context.Context) (*dto.ReorgsResponse, error)
}

type reorgService struct {
	repo repository.ReorgRepo
}

func NewReorgService(repo repository.ReorgRepo) ReorgService {
	return &reorgService{repo: repo}
}

func (s *reorgService) GetReorgs(ctx context.Context) (*dto.ReorgsResponse, error) {
	return s.repo.GetReorgs(ctx)
}
