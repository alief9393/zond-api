package service

import (
	"context"
	"zond-api/internal/api/dto"
	chainRepo "zond-api/internal/api/repository/chain"
)

type ChainService interface {
	GetChainInfo(ctx context.Context) (*dto.ChainResponse, error)
}

type chainService struct {
	repo chainRepo.ChainRepo
}

func NewChainService(repo chainRepo.ChainRepo) ChainService {
	return &chainService{repo: repo}
}

func (s *chainService) GetChainInfo(ctx context.Context) (*dto.ChainResponse, error) {
	return s.repo.GetChainInfo(ctx)
}
