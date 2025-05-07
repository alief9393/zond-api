package repository

import (
	"context"
	"fmt"

	"zond-api/internal/api/dto"
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

func (r *TransactionRepoPG) GetTransactionsByBlockNumber(blockNumber int64) ([]model.Transaction, error) {
	rows, err := r.db.Query(context.Background(), `
		SELECT tx_hash, block_number, from_address, to_address, value, gas, gas_price,
		       type, chain_id, access_list, max_fee_per_gas, max_priority_fee_per_gas,
		       transaction_index, cumulative_gas_used, is_successful, retrieved_from,
		       is_canonical
		FROM transactions
		WHERE block_number = $1
		ORDER BY transaction_index ASC`, blockNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []model.Transaction
	for rows.Next() {
		var tx model.Transaction
		err := rows.Scan(
			&tx.TxHash, &tx.BlockNumber, &tx.FromAddress, &tx.ToAddress,
			&tx.Value, &tx.Gas, &tx.GasPrice, &tx.Type, &tx.ChainID, &tx.AccessList,
			&tx.MaxFeePerGas, &tx.MaxPriorityFeePerGas, &tx.TransactionIndex,
			&tx.CumulativeGasUsed, &tx.IsSuccessful, &tx.RetrievedFrom, &tx.IsCanonical,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, tx)
	}
	return transactions, nil
}

func (r *TransactionRepoPG) GetTransactionMetrics(ctx context.Context) (*dto.TransactionMetricsResponse, error) {
	var metrics dto.TransactionMetricsResponse
	err := r.db.QueryRow(ctx, `
		SELECT
			(SELECT COUNT(*) FROM transactions WHERE timestamp >= NOW() - INTERVAL '24 HOURS') as transactions_24h,
			(SELECT COUNT(*) FROM transactions WHERE is_pending = TRUE AND timestamp >= NOW() - INTERVAL '1 HOUR') as pending_transactions_1h,
			(SELECT COALESCE(SUM(fee_eth), 0) FROM transactions WHERE timestamp >= NOW() - INTERVAL '24 HOURS') as network_fee_eth_24h,
			(SELECT COALESCE(AVG(fee_usd), 0) FROM transactions WHERE timestamp >= NOW() - INTERVAL '24 HOURS') as avg_fee_usd_24h
	`).Scan(
		&metrics.Transactions24h,
		&metrics.PendingTransactions1h,
		&metrics.NetworkFeeETH24h,
		&metrics.AvgFeeUSD24h,
	)
	if err != nil {
		return nil, err
	}
	return &metrics, nil
}

func (r *TransactionRepoPG) CountTransactionsWithFilter(ctx context.Context, method, from, to string) (int, error) {
	query := `SELECT COUNT(*) FROM transactions WHERE 1=1`
	args := []interface{}{}
	argPos := 1

	if method != "" {
		query += fmt.Sprintf(" AND method = $%d", argPos)
		args = append(args, method)
		argPos++
	}
	if from != "" {
		query += fmt.Sprintf(" AND encode(from_address, 'hex') ILIKE $%d", argPos)
		args = append(args, from+"%")
		argPos++
	}
	if to != "" {
		query += fmt.Sprintf(" AND encode(to_address, 'hex') ILIKE $%d", argPos)
		args = append(args, to+"%")
		argPos++
	}

	var total int
	err := r.db.QueryRow(ctx, query, args...).Scan(&total)
	return total, err
}

func (r *TransactionRepoPG) GetLatestTransactionsWithFilter(ctx context.Context, page, limit int, method, from, to string) ([]model.Transaction, error) {
	query := `SELECT tx_hash, block_number, from_address, to_address, value, gas, gas_price, type,
	                 chain_id, access_list, max_fee_per_gas, max_priority_fee_per_gas, transaction_index,
	                 cumulative_gas_used, is_successful, retrieved_from, is_canonical
	          FROM transactions WHERE 1=1`
	args := []interface{}{}
	argPos := 1

	if method != "" {
		query += fmt.Sprintf(" AND method = $%d", argPos)
		args = append(args, method)
		argPos++
	}
	if from != "" {
		query += fmt.Sprintf(" AND encode(from_address, 'hex') ILIKE $%d", argPos)
		args = append(args, from+"%")
		argPos++
	}
	if to != "" {
		query += fmt.Sprintf(" AND encode(to_address, 'hex') ILIKE $%d", argPos)
		args = append(args, to+"%")
		argPos++
	}

	query += fmt.Sprintf(" ORDER BY block_number DESC, transaction_index ASC LIMIT $%d OFFSET $%d", argPos, argPos+1)
	args = append(args, limit, (page-1)*limit)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var txs []model.Transaction
	for rows.Next() {
		var tx model.Transaction
		if err := rows.Scan(
			&tx.TxHash, &tx.BlockNumber, &tx.FromAddress, &tx.ToAddress, &tx.Value, &tx.Gas, &tx.GasPrice,
			&tx.Type, &tx.ChainID, &tx.AccessList, &tx.MaxFeePerGas, &tx.MaxPriorityFeePerGas,
			&tx.TransactionIndex, &tx.CumulativeGasUsed, &tx.IsSuccessful, &tx.RetrievedFrom, &tx.IsCanonical,
		); err != nil {
			return nil, err
		}
		txs = append(txs, tx)
	}

	return txs, nil
}

func (r *TransactionRepoPG) CountPendingTransactions(ctx context.Context, method, from, to string) (int, error) {
	var count int
	query := `
		SELECT COUNT(*) FROM transactions
		WHERE block_number IS NULL
	`

	args := []interface{}{}
	i := 1

	if method != "" {
		query += fmt.Sprintf(" AND type = $%d", i)
		args = append(args, method)
		i++
	}
	if from != "" {
		query += fmt.Sprintf(" AND encode(from_address, 'hex') ILIKE $%d", i)
		args = append(args, from+"%")
		i++
	}
	if to != "" {
		query += fmt.Sprintf(" AND encode(to_address, 'hex') ILIKE $%d", i)
		args = append(args, to+"%")
		i++
	}

	err := r.db.QueryRow(ctx, query, args...).Scan(&count)
	return count, err
}

func (r *TransactionRepoPG) GetPendingTransactions(ctx context.Context, page, limit int, method, from, to string) ([]model.Transaction, error) {
	var transactions []model.Transaction

	offset := (page - 1) * limit
	query := `
		SELECT tx_hash, block_number, from_address, to_address, value, gas, gas_price, type,
		       chain_id, access_list, max_fee_per_gas, max_priority_fee_per_gas, transaction_index,
		       cumulative_gas_used, is_successful, retrieved_from, is_canonical
		FROM transactions
		WHERE block_number IS NULL
	`

	args := []interface{}{}
	i := 1

	if method != "" {
		query += fmt.Sprintf(" AND type = $%d", i)
		args = append(args, method)
		i++
	}
	if from != "" {
		query += fmt.Sprintf(" AND encode(from_address, 'hex') ILIKE $%d", i)
		args = append(args, from+"%")
		i++
	}
	if to != "" {
		query += fmt.Sprintf(" AND encode(to_address, 'hex') ILIKE $%d", i)
		args = append(args, to+"%")
		i++
	}

	query += fmt.Sprintf(" ORDER BY transaction_index DESC LIMIT $%d OFFSET $%d", i, i+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(ctx, query, args...)
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

func (r *TransactionRepoPG) GetContractTransactions(ctx context.Context, page, limit int, method, from, to string) ([]model.Transaction, error) {
	offset := (page - 1) * limit
	var transactions []model.Transaction

	query := `
		SELECT tx_hash, block_number, from_address, to_address, value, gas, gas_price, type,
		       chain_id, access_list, max_fee_per_gas, max_priority_fee_per_gas, transaction_index,
		       cumulative_gas_used, is_successful, retrieved_from, is_canonical
		FROM transactions
		WHERE to_address IS NOT NULL
		AND is_contract = TRUE
	`

	args := []interface{}{}
	argPos := 1

	if method != "" {
		query += fmt.Sprintf(" AND method = $%d", argPos)
		args = append(args, method)
		argPos++
	}

	if from != "" {
		query += fmt.Sprintf(" AND encode(from_address, 'hex') = $%d", argPos)
		args = append(args, from)
		argPos++
	}

	if to != "" {
		query += fmt.Sprintf(" AND encode(to_address, 'hex') = $%d", argPos)
		args = append(args, to)
		argPos++
	}

	query += fmt.Sprintf(" ORDER BY block_number DESC OFFSET $%d LIMIT $%d", argPos, argPos+1)
	args = append(args, offset, limit)

	rows, err := r.db.Query(ctx, query, args...)
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

func (r *TransactionRepoPG) CountContractTransactions(ctx context.Context, method, from, to string) (int, error) {
	query := `
		SELECT COUNT(*)
		FROM transactions
		WHERE to_address IS NOT NULL
		AND is_contract = TRUE
	`

	args := []interface{}{}
	argPos := 1

	if method != "" {
		query += fmt.Sprintf(" AND method = $%d", argPos)
		args = append(args, method)
		argPos++
	}

	if from != "" {
		query += fmt.Sprintf(" AND encode(from_address, 'hex') = $%d", argPos)
		args = append(args, from)
		argPos++
	}

	if to != "" {
		query += fmt.Sprintf(" AND encode(to_address, 'hex') = $%d", argPos)
		args = append(args, to)
		argPos++
	}

	var count int
	err := r.db.QueryRow(ctx, query, args...).Scan(&count)
	return count, err
}

func (r *TransactionRepoPG) GetDailyTransactionStats(days int) ([]dto.DailyTransactionStat, error) {
	rows, err := r.db.Query(context.Background(), `
		SELECT to_char(block_timestamp::date, 'YYYY-MM-DD') AS date, COUNT(*) AS count
		FROM transactions
		WHERE block_timestamp >= current_date - $1::int
		GROUP BY date
		ORDER BY date DESC
	`, days)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []dto.DailyTransactionStat
	for rows.Next() {
		var stat dto.DailyTransactionStat
		if err := rows.Scan(&stat.Date, &stat.Count); err != nil {
			return nil, err
		}
		stats = append(stats, stat)
	}

	return stats, nil
}

func (r *TransactionRepoPG) GetAverageTPS(lastNBlocks int) (float64, error) {
	row := r.db.QueryRow(context.Background(), `
		WITH recent_blocks AS (
			SELECT block_number, block_timestamp, transaction_count
			FROM blocks
			ORDER BY block_number DESC
			LIMIT $1
		)
		SELECT 
			COALESCE(SUM(transaction_count)::float / EXTRACT(EPOCH FROM MAX(block_timestamp) - MIN(block_timestamp)), 0)
		FROM recent_blocks
		WHERE MAX(block_timestamp) > MIN(block_timestamp)
	`, lastNBlocks)

	var tps float64
	err := row.Scan(&tps)
	if err != nil {
		return 0, err
	}
	return tps, nil
}

func (r *TransactionRepoPG) GetDailyFeeStats(days int) ([]dto.DailyFeeStat, error) {
	rows, err := r.db.Query(context.Background(), `
		SELECT 
			TO_CHAR(block_timestamp::date, 'YYYY-MM-DD') AS date,
			SUM(fee_eth) AS total_eth,
			AVG(fee_usd) AS avg_usd
		FROM transactions
		WHERE block_timestamp >= current_date - $1::int
		GROUP BY date
		ORDER BY date ASC
	`, days)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []dto.DailyFeeStat
	for rows.Next() {
		var stat dto.DailyFeeStat
		if err := rows.Scan(&stat.Date, &stat.NetworkFeeETH, &stat.AvgFeeUSD); err != nil {
			return nil, err
		}
		stats = append(stats, stat)
	}
	return stats, nil
}
