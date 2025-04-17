package service

import (
	"context"
	"zond-api/internal/api/dto"
	"zond-api/internal/api/repository"
)

type ForkService interface {
	GetForks(ctx context.Context) (*dto.ForksResponse, error)
}

type forkService struct {
	repo repository.ForkRepo
}

func NewForkService(repo repository.ForkRepo) ForkService {
	return &forkService{repo: repo}
}

func (s *forkService) GetForks(ctx context.Context) (*dto.ForksResponse, error) {
	return s.repo.GetForks(ctx)
}
