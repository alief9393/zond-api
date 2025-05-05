package service

import (
	"context"
	"zond-api/internal/api/dto"
	validatorRepo "zond-api/internal/api/repository/validator"
)

type ValidatorService interface {
	GetValidators(ctx context.Context) (*dto.ValidatorsResponse, error)
	GetValidatorByID(ctx context.Context, index int) (*dto.ValidatorDetailResponse, error)
}

type validatorService struct {
	repo validatorRepo.ValidatorRepo
}

func NewValidatorService(repo validatorRepo.ValidatorRepo) ValidatorService {
	return &validatorService{repo: repo}
}

func (s *validatorService) GetValidators(ctx context.Context) (*dto.ValidatorsResponse, error) {
	return s.repo.GetValidators(ctx)
}

func (s *validatorService) GetValidatorByID(ctx context.Context, index int) (*dto.ValidatorDetailResponse, error) {
	return s.repo.GetValidatorByID(ctx, index)
}
