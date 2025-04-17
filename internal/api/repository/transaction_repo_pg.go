package repository

import (
	"context"

	"zond-api/internal/domain/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionRepoPG struct {
	db *pgxpool.Pool
}

func NewTransactionRepoPG(db *pgxpool.Pool) *TransactionRepoPG {
	return &TransactionRepoPG{db: db}
}

func (r *TransactionRepoPG) GetLatestTransactions(limit, offset int) ([]model.Transaction, error) {
	var transactions []model.Transaction
	rows, err := r.db.Query(context.Background(), `
        SELECT tx_hash, block_number, from_address, to_address, value, gas, gas_price, type,
               chain_id, access_list, max_fee_per_gas, max_priority_fee_per_gas, transaction_index,
               cumulative_gas_used, is_successful, retrieved_from, is_canonical
        FROM Transactions
        WHERE is_canonical = TRUE
        ORDER BY block_number DESC, transaction_index DESC
        LIMIT $1 OFFSET $2
    `, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tx model.Transaction
		err := rows.Scan(
			&tx.TxHash, &tx.BlockNumber, &tx.FromAddress, &tx.ToAddress, &tx.Value, &tx.Gas,
			&tx.GasPrice, &tx.Type, &tx.ChainID, &tx.AccessList, &tx.MaxFeePerGas,
			&tx.MaxPriorityFeePerGas, &tx.TransactionIndex, &tx.CumulativeGasUsed,
			&tx.IsSuccessful, &tx.RetrievedFrom, &tx.IsCanonical,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, tx)
	}
	return transactions, nil
}

func (r *TransactionRepoPG) GetTransactionByHash(txHash []byte) (*model.Transaction, error) {
	var tx model.Transaction
	err := r.db.QueryRow(context.Background(), `
        SELECT tx_hash, block_number, from_address, to_address, value, gas, gas_price, type,
               chain_id, access_list, max_fee_per_gas, max_priority_fee_per_gas, transaction_index,
               cumulative_gas_used, is_successful, retrieved_from, is_canonical
        FROM Transactions
        WHERE tx_hash = $1
    `, txHash).Scan(
		&tx.TxHash, &tx.BlockNumber, &tx.FromAddress, &tx.ToAddress, &tx.Value, &tx.Gas,
		&tx.GasPrice, &tx.Type, &tx.ChainID, &tx.AccessList, &tx.MaxFeePerGas,
		&tx.MaxPriorityFeePerGas, &tx.TransactionIndex, &tx.CumulativeGasUsed,
		&tx.IsSuccessful, &tx.RetrievedFrom, &tx.IsCanonical,
	)
	if err != nil {
		return nil, err
	}
	return &tx, nil
}
