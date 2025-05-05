package repository

import (
	"context"
	"zond-api/internal/api/dto"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AddressRepo interface {
	GetAddressBalance(ctx context.Context, address string) (*dto.AddressResponse, error)
	GetAddressTransactions(ctx context.Context, address string) (*dto.TransactionsResponse, error)
	GetTopAddresses(ctx context.Context, limit int) (*dto.TopAddressesResponse, error)
	GetAddressDetails(ctx context.Context, address string) (*dto.AddressDetailResponse, error)
}

type AddressRepoPG struct {
	db *pgxpool.Pool
}

func NewAddressRepoPG(db *pgxpool.Pool) AddressRepo {
	return &AddressRepoPG{db: db}
}

func (r *AddressRepoPG) GetAddressBalance(ctx context.Context, address string) (*dto.AddressResponse, error) {
	var addr dto.AddressResponse
	err := r.db.QueryRow(ctx, "SELECT address, balance FROM addresses WHERE address = $1", address).
		Scan(&addr.Address, &addr.Balance)
	if err != nil {
		return nil, err
	}
	return &addr, nil
}

func (r *AddressRepoPG) GetAddressTransactions(ctx context.Context, address string) (*dto.TransactionsResponse, error) {
	rows, err := r.db.Query(ctx, `
		SELECT tx_hash, block_number, from_address, to_address, value, gas, gas_price, type, chain_id, 
		       access_list, max_fee_per_gas, max_priority_fee_per_gas, transaction_index, cumulative_gas_used, 
		       is_successful, retrieved_from, is_canonical
		FROM transactions 
		WHERE from_address = $1 OR to_address = $1`, address)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []dto.TransactionResponse
	for rows.Next() {
		var tx dto.TransactionResponse
		err := rows.Scan(
			&tx.TxHash, &tx.BlockNumber, &tx.FromAddress, &tx.ToAddress, &tx.Value, &tx.Gas, &tx.GasPrice,
			&tx.Type, &tx.ChainID, &tx.AccessList, &tx.MaxFeePerGas, &tx.MaxPriorityFeePerGas,
			&tx.TransactionIndex, &tx.CumulativeGasUsed, &tx.IsSuccessful, &tx.RetrievedFrom, &tx.IsCanonical,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, tx)
	}

	return &dto.TransactionsResponse{Transactions: transactions}, nil
}

func (r *AddressRepoPG) GetTopAddresses(ctx context.Context, limit int) (*dto.TopAddressesResponse, error) {
	rows, err := r.db.Query(ctx, `
		SELECT address, balance 
		FROM addresses 
		ORDER BY balance::numeric DESC 
		LIMIT $1`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addresses []dto.AddressResponse
	for rows.Next() {
		var addr dto.AddressResponse
		if err := rows.Scan(&addr.Address, &addr.Balance); err != nil {
			return nil, err
		}
		addresses = append(addresses, addr)
	}
	return &dto.TopAddressesResponse{Addresses: addresses}, nil
}

func (r *AddressRepoPG) GetAddressDetails(ctx context.Context, address string) (*dto.AddressDetailResponse, error) {
	var result dto.AddressDetailResponse
	result.Address = address

	err := r.db.QueryRow(ctx, `
		SELECT COALESCE(SUM(value), 0), COUNT(*), COALESCE(SUM(gas), 0)
		FROM transactions
		WHERE from_address = decode($1, 'hex') OR to_address = decode($1, 'hex')
	`, address[2:]).Scan(&result.Balance, &result.TransactionCount, &result.TotalGasUsed)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
