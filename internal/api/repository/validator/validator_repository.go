package repository

import (
	"context"
	"zond-api/internal/api/dto"
)

type ValidatorRepo interface {
	GetValidators(ctx context.Context) (*dto.ValidatorsResponse, error)
	GetValidatorByID(ctx context.Context, index int) (*dto.ValidatorDetailResponse, error)
}
