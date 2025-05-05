package repository

import (
	"context"
	"zond-api/internal/api/dto"
)

type ReorgRepository interface {
	GetReorgs(ctx context.Context) (*dto.ReorgsResponse, error)
}
