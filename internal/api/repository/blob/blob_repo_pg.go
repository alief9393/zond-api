package blob

import (
	"context"
	"fmt"

	"zond-api/internal/domain/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BlobRepoPG struct {
	db *pgxpool.Pool
}

func NewBlobRepoPG(db *pgxpool.Pool) *BlobRepoPG {
	return &BlobRepoPG{db: db}
}

func (r *BlobRepoPG) GetBlobs(limit, offset int) ([]model.Blob, int, error) {
	query := `
		SELECT
			versioned_hash,
			tx_hash,
			block_number,
			timestamp,
			blob_sender,
			gas_price,
			size,
			retrieved_from
		FROM blobs
		ORDER BY block_number DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(context.Background(), query, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var blobs []model.Blob
	for rows.Next() {
		var b model.Blob
		if err := rows.Scan(
			&b.VersionedHash,
			&b.TxHash,
			&b.BlockNumber,
			&b.Timestamp,
			&b.BlobSender,
			&b.GasPrice,
			&b.Size,
			&b.RetrievedFrom,
		); err != nil {
			return nil, 0, fmt.Errorf("scan error: %w", err)
		}
		blobs = append(blobs, b)
	}

	var total int
	err = r.db.QueryRow(context.Background(), `SELECT COUNT(*) FROM blobs`).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("count error: %w", err)
	}

	return blobs, total, nil
}
