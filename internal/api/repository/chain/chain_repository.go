package repository

import (
	"context"
	"zond-api/internal/api/dto"
)

type ChainRepository interface {
	GetChainInfo(ctx context.Context) (*dto.ChainResponse, error)
}
