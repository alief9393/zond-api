package service

import (
	"context"
	"zond-api/internal/api/dto"
	forkRepo "zond-api/internal/api/repository/fork"
)

type ForkService interface {
	GetForks(ctx context.Context) (*dto.ForksResponse, error)
}

type forkService struct {
	repo forkRepo.ForkRepo
}

func NewForkService(repo forkRepo.ForkRepo) ForkService {
	return &forkService{repo: repo}
}

func (s *forkService) GetForks(ctx context.Context) (*dto.ForksResponse, error) {
	return s.repo.GetForks(ctx)
}
