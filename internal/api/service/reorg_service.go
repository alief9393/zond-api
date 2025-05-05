package service

import (
	"context"
	"zond-api/internal/api/dto"
	reorgRepo "zond-api/internal/api/repository/reorg"
)

type ReorgService interface {
	GetReorgs(ctx context.Context) (*dto.ReorgsResponse, error)
}

type reorgService struct {
	repo reorgRepo.ReorgRepo
}

func NewReorgService(repo reorgRepo.ReorgRepo) ReorgService {
	return &reorgService{repo: repo}
}

func (s *reorgService) GetReorgs(ctx context.Context) (*dto.ReorgsResponse, error) {
	return s.repo.GetReorgs(ctx)
}
