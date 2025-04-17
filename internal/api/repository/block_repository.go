package repository

import "zond-api/internal/domain/model"

type BlockRepository interface {
	GetLatestBlocks(limit, offset int) ([]model.Block, error)
	GetBlockByNumber(blockNumber int64) (*model.Block, error)
	GetForkedBlocks(limit, offset int) ([]model.Block, error)
}
