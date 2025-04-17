package service

import (
	"context"
	"zond-api/internal/api/dto"
	"zond-api/internal/api/repository"
)

type ChainService interface {
	GetChainInfo(ctx context.Context) (*dto.ChainResponse, error)
}

type chainService struct {
	repo repository.ChainRepo
}

func NewChainService(repo repository.ChainRepo) ChainService {
	return &chainService{repo: repo}
}

func (s *chainService) GetChainInfo(ctx context.Context) (*dto.ChainResponse, error) {
	return s.repo.GetChainInfo(ctx)
}
