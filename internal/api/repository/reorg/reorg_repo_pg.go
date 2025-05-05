package repository

import (
	"context"
	"zond-api/internal/api/dto"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ReorgRepo interface {
	GetReorgs(ctx context.Context) (*dto.ReorgsResponse, error)
}

type ReorgRepoPG struct {
	db *pgxpool.Pool
}

func NewReorgRepoPG(db *pgxpool.Pool) ReorgRepo {
	return &ReorgRepoPG{db: db}
}

func (r *ReorgRepoPG) GetReorgs(ctx context.Context) (*dto.ReorgsResponse, error) {
	rows, err := r.db.Query(ctx, `
		SELECT block_number, old_block_hash, new_block_hash, depth, timestamp, chain_id, retrieved_from
		FROM reorgs
		ORDER BY timestamp DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reorgs []dto.ReorgResponse
	for rows.Next() {
		var reorg dto.ReorgResponse
		err := rows.Scan(
			&reorg.BlockNumber, &reorg.OldBlockHash, &reorg.NewBlockHash,
			&reorg.Depth, &reorg.Timestamp, &reorg.ChainID, &reorg.RetrievedFrom,
		)
		if err != nil {
			return nil, err
		}
		reorgs = append(reorgs, reorg)
	}

	return &dto.ReorgsResponse{Reorgs: reorgs}, nil
}
