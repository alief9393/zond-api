package repository

import (
	"context"
	"zond-api/internal/api/dto"
)

type ForkRepository interface {
	GetForks(ctx context.Context) (*dto.ForksResponse, error)
}
