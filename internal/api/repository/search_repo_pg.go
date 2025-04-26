package repository

import (
	"context"

	"zond-api/internal/api/dto"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SearchRepository interface {
	Search(ctx context.Context, query string) ([]dto.SearchResult, error)
}

type searchRepoPG struct {
	db *pgxpool.Pool
}

func NewSearchRepoPG(db *pgxpool.Pool) SearchRepository {
	return &searchRepoPG{db: db}
}

func (r *searchRepoPG) Search(ctx context.Context, query string) ([]dto.SearchResult, error) {
	var results []dto.SearchResult

	// Search blocks by hash
	var block dto.BlockResponse
	err := r.db.QueryRow(ctx, `
		SELECT block_number, block_hash, timestamp, miner_address, canonical, 
		       parent_hash, gas_used, gas_limit, size, transaction_count, 
		       extra_data, base_fee_per_gas, transactions_root, state_root, 
		       receipts_root, logs_bloom, chain_id, retrieved_from
		FROM blocks WHERE block_hash = $1`, query).
		Scan(&block.BlockNumber, &block.BlockHash, &block.Timestamp, &block.MinerAddress,
			&block.Canonical, &block.ParentHash, &block.GasUsed, &block.GasLimit,
			&block.Size, &block.TransactionCount, &block.ExtraData, &block.BaseFeePerGas,
			&block.TransactionsRoot, &block.StateRoot, &block.ReceiptsRoot,
			&block.LogsBloom, &block.ChainID, &block.RetrievedFrom)
	if err == nil {
		results = append(results, dto.SearchResult{Type: "block", Value: block})
	}

	// Search transactions by hash
	var tx dto.TransactionDTO
	err = r.db.QueryRow(ctx, `
		SELECT hash, block_number, from_address, to_address, value 
		FROM transactions WHERE hash = $1`, query).
		Scan(&tx.Hash, &tx.BlockNumber, &tx.From, &tx.To, &tx.Value)
	if err == nil {
		results = append(results, dto.SearchResult{Type: "transaction", Value: tx})
	}

	// Search addresses
	var addr dto.AddressResponse
	err = r.db.QueryRow(ctx, `
		SELECT address, balance 
		FROM addresses WHERE address = $1`, query).
		Scan(&addr.Address, &addr.Balance)
	if err == nil {
		results = append(results, dto.SearchResult{Type: "address", Value: addr})
	}

	return results, nil
}
