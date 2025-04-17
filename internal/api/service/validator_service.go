package service

import (
	"context"
	"zond-api/internal/api/dto"
	"zond-api/internal/api/repository"
)

type ValidatorService interface {
	GetValidators(ctx context.Context) (*dto.ValidatorsResponse, error)
}

type validatorService struct {
	repo repository.ValidatorRepo
}

func NewValidatorService(repo repository.ValidatorRepo) ValidatorService {
	return &validatorService{repo: repo}
}

func (s *validatorService) GetValidators(ctx context.Context) (*dto.ValidatorsResponse, error) {
	return s.repo.GetValidators(ctx)
}
