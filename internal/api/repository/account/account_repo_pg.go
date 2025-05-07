package account

import (
	"context"
	"zond-api/internal/domain/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountRepoPG struct {
	db *pgxpool.Pool
}

func NewAccountRepoPG(db *pgxpool.Pool) *AccountRepoPG {
	return &AccountRepoPG{db: db}
}

func (r *AccountRepoPG) GetTopAccounts(ctx context.Context, page, limit int) ([]model.Account, int, error) {
	offset := (page - 1) * limit
	rows, err := r.db.Query(ctx, `
		SELECT address, name_tag, balance, percentage, tx_count
		FROM accounts
		ORDER BY balance DESC
		LIMIT $1 OFFSET $2
	`, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var accounts []model.Account
	for rows.Next() {
		var acc model.Account
		if err := rows.Scan(&acc.Address, &acc.NameTag, &acc.Balance, &acc.Percentage, &acc.TxCount); err != nil {
			return nil, 0, err
		}
		accounts = append(accounts, acc)
	}

	var total int
	err = r.db.QueryRow(ctx, `SELECT COUNT(*) FROM accounts`).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return accounts, total, nil
}
