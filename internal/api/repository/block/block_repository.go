package repository

import (
	"context"
	"zond-api/internal/domain/model"
)

type BlockRepository interface {
	GetLatestBlocks(limit, offset int) ([]model.Block, error)
	GetBlockByNumber(blockNumber int64) (*model.Block, error)
	GetForkedBlocks(limit, offset int) ([]model.Block, error)
	GetBlockByHash(ctx context.Context, hash string) (*model.Block, error)
	GetPaginatedBlocks(ctx context.Context, page, limit int) ([]model.Block, int, error)
}
