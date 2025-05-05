package repository

import (
	"context"
	"zond-api/internal/api/dto"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ForkRepo interface {
	GetForks(ctx context.Context) (*dto.ForksResponse, error)
}

type ForkRepoPG struct {
	db *pgxpool.Pool
}

func NewForkRepoPG(db *pgxpool.Pool) ForkRepo {
	return &ForkRepoPG{db: db}
}

func (r *ForkRepoPG) GetForks(ctx context.Context) (*dto.ForksResponse, error) {
	rows, err := r.db.Query(ctx, `
		SELECT fork_name, block_number, timestamp, chain_id, retrieved_from
		FROM forks
		ORDER BY block_number DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var forks []dto.ForkResponse
	for rows.Next() {
		var fork dto.ForkResponse
		err := rows.Scan(
			&fork.ForkName, &fork.BlockNumber, &fork.Timestamp,
			&fork.ChainID, &fork.RetrievedFrom,
		)
		if err != nil {
			return nil, err
		}
		forks = append(forks, fork)
	}

	return &dto.ForksResponse{Forks: forks}, nil
}
