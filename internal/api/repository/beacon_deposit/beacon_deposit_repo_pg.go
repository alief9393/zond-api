package beacon_deposit

import (
	"context"
	"zond-api/internal/domain/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type beaconDepositRepoPG struct {
	db *pgxpool.Pool
}

func NewBeaconDepositRepoPG(db *pgxpool.Pool) BeaconDepositRepository {
	return &beaconDepositRepoPG{db: db}
}

func (r *beaconDepositRepoPG) GetBeaconDeposits(ctx context.Context, limit, offset int) ([]model.BeaconDeposit, error) {
	rows, err := r.db.Query(ctx, `
		SELECT block_number, tx_hash, amount, validator_index, timestamp, log_index, retrieved_from
		FROM beacon_deposits
		ORDER BY block_number DESC
		LIMIT $1 OFFSET $2
	`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var deposits []model.BeaconDeposit
	for rows.Next() {
		var d model.BeaconDeposit
		err := rows.Scan(
			&d.BlockNumber, &d.TxHash, &d.Amount, &d.ValidatorIndex,
			&d.Timestamp, &d.LogIndex, &d.RetrievedFrom,
		)
		if err != nil {
			return nil, err
		}
		deposits = append(deposits, d)
	}
	return deposits, nil
}

func (r *beaconDepositRepoPG) CountBeaconDeposits(ctx context.Context) (int, error) {
	var count int
	err := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM beacon_deposits`).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
