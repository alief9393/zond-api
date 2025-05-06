package beacon_withdrawal

import (
	"context"
	"zond-api/internal/domain/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type beaconWithdrawalRepoPG struct {
	db *pgxpool.Pool
}

func NewBeaconWithdrawalRepoPG(db *pgxpool.Pool) BeaconWithdrawalRepository {
	return &beaconWithdrawalRepoPG{db: db}
}

func (r *beaconWithdrawalRepoPG) GetBeaconWithdrawals(ctx context.Context, limit, offset int) ([]model.BeaconWithdrawal, error) {
	rows, err := r.db.Query(ctx, `
		SELECT block_number, tx_hash, validator_index, amount, timestamp, retrieved_from, log_index
		FROM beacon_withdrawals
		ORDER BY block_number DESC
		LIMIT $1 OFFSET $2
	`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var withdrawals []model.BeaconWithdrawal
	for rows.Next() {
		var w model.BeaconWithdrawal
		err := rows.Scan(
			&w.BlockNumber,
			&w.TxHash,
			&w.ValidatorIndex,
			&w.Amount,
			&w.Timestamp,
			&w.RetrievedFrom,
			&w.LogIndex,
		)
		if err != nil {
			return nil, err
		}
		withdrawals = append(withdrawals, w)
	}
	return withdrawals, nil
}

func (r *beaconWithdrawalRepoPG) CountBeaconWithdrawals(ctx context.Context) (int, error) {
	var count int
	err := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM beacon_withdrawals`).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
