package account_abstraction

import (
	"context"
	"fmt"

	"zond-api/internal/domain/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountAbstractionRepoPG struct {
	db *pgxpool.Pool
}

func NewAccountAbstractionRepoPG(db *pgxpool.Pool) *AccountAbstractionRepoPG {
	return &AccountAbstractionRepoPG{db: db}
}

func (r *AccountAbstractionRepoPG) GetBundleTransactions(limit, offset int) ([]model.BundleTransaction, int, error) {
	rows, err := r.db.Query(context.Background(), `
		SELECT bundle_txn_hash, block_number, timestamp, bundler, entry_point,
		       aa_txns_count, amount, gas_price
		FROM bundle_transactions
		ORDER BY block_number DESC
		LIMIT $1 OFFSET $2
	`, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("query bundles error: %w", err)
	}
	defer rows.Close()

	var bundles []model.BundleTransaction
	for rows.Next() {
		var b model.BundleTransaction
		if err := rows.Scan(&b.Hash, &b.BlockNumber, &b.Timestamp, &b.Bundler, &b.EntryPoint,
			&b.AATxnCount, &b.Amount, &b.GasPrice); err != nil {
			return nil, 0, fmt.Errorf("scan bundle error: %w", err)
		}
		bundles = append(bundles, b)
	}

	var total int
	err = r.db.QueryRow(context.Background(), `SELECT COUNT(*) FROM bundle_transactions`).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("count bundle error: %w", err)
	}

	return bundles, total, nil
}

func (r *AccountAbstractionRepoPG) GetAATransactions(limit, offset int) ([]model.AATransaction, int, error) {
	rows, err := r.db.Query(context.Background(), `
		SELECT aa_txn_hash, bundle_txn_hash, method, block_number, timestamp,
		       from_address, bundler, entry_point, gas_price
		FROM aa_transactions
		ORDER BY block_number DESC
		LIMIT $1 OFFSET $2
	`, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("query aa error: %w", err)
	}
	defer rows.Close()

	var txns []model.AATransaction
	for rows.Next() {
		var a model.AATransaction
		if err := rows.Scan(&a.Hash, &a.BundleHash, &a.Method, &a.BlockNumber, &a.Timestamp,
			&a.From, &a.Bundler, &a.EntryPoint, &a.GasPrice); err != nil {
			return nil, 0, fmt.Errorf("scan aa error: %w", err)
		}
		txns = append(txns, a)
	}

	var total int
	err = r.db.QueryRow(context.Background(), `SELECT COUNT(*) FROM aa_transactions`).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("count aa error: %w", err)
	}

	return txns, total, nil
}
