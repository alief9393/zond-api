package repository

import (
	"context"
	"zond-api/internal/api/dto"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ChainRepo interface {
	GetChainInfo(ctx context.Context) (*dto.ChainResponse, error)
}

type ChainRepoPG struct {
	db *pgxpool.Pool
}

func NewChainRepoPG(db *pgxpool.Pool) ChainRepo {
	return &ChainRepoPG{db: db}
}

func (r *ChainRepoPG) GetChainInfo(ctx context.Context) (*dto.ChainResponse, error) {
	var chain dto.ChainResponse

	// Get the latest block number
	err := r.db.QueryRow(ctx, `
		SELECT block_number, retrieved_from
		FROM blocks
		ORDER BY block_number DESC
		LIMIT 1`).
		Scan(&chain.LatestBlock, &chain.RetrievedFrom)
	if err != nil {
		return nil, err
	}

	// Get the total number of blocks
	err = r.db.QueryRow(ctx, `SELECT COUNT(*) FROM blocks`).
		Scan(&chain.TotalBlocks)
	if err != nil {
		return nil, err
	}

	// Hardcode chain ID for now (replace with actual value or fetch from config)
	chain.ChainID = 1337 // Example chain ID for Zond testnet

	return &chain, nil
}
