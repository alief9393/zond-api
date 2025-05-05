package repository

import (
	"context"
	"fmt"
	"zond-api/internal/api/dto"

	"github.com/jackc/pgx/v5/pgxpool"
)

type searchRepoPG struct {
	db *pgxpool.Pool
}

func NewSearchRepoPG(db *pgxpool.Pool) SearchRepository {
	return &searchRepoPG{db: db}
}

func (r *searchRepoPG) Search(ctx context.Context, query string) ([]dto.SearchResult, error) {
	var results []dto.SearchResult
	// ... implementation ...
	return results, nil
}

func (r *searchRepoPG) GetSuggestions(ctx context.Context, query string) ([]dto.Suggestion, error) {
	var suggestions []dto.Suggestion

	// Address match
	rows, err := r.db.Query(ctx, `
		SELECT encode(from_address, 'hex') 
		FROM transactions
		WHERE encode(from_address, 'hex') ILIKE $1
		LIMIT 5
	`, query+"%")
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var addr string
			_ = rows.Scan(&addr)
			suggestions = append(suggestions, dto.Suggestion{Type: "address", Value: "0x" + addr})
		}
	}

	// TX match
	row := r.db.QueryRow(ctx, `
		SELECT encode(tx_hash, 'hex')
		FROM transactions
		WHERE encode(tx_hash, 'hex') ILIKE $1
		LIMIT 1
	`, query+"%")

	var tx string
	if err := row.Scan(&tx); err == nil {
		suggestions = append(suggestions, dto.Suggestion{Type: "transaction", Value: "0x" + tx})
	}

	// Block number match
	row = r.db.QueryRow(ctx, `
		SELECT block_number
		FROM blocks
		WHERE CAST(block_number AS TEXT) ILIKE $1
		LIMIT 1
	`, query+"%")

	var blk int64
	if err := row.Scan(&blk); err == nil {
		suggestions = append(suggestions, dto.Suggestion{Type: "block", Value: fmt.Sprintf("%d", blk)})
	}

	return suggestions, nil
}
