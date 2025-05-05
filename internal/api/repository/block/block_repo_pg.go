package repository

import (
	"context"

	"zond-api/internal/domain/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BlockRepoPG struct {
	db *pgxpool.Pool
}

func NewBlockRepoPG(db *pgxpool.Pool) *BlockRepoPG {
	return &BlockRepoPG{db: db}
}

func (r *BlockRepoPG) GetLatestBlocks(limit, offset int) ([]model.Block, error) {
	var blocks []model.Block
	rows, err := r.db.Query(context.Background(), `
        SELECT block_number, block_hash, timestamp, miner_address, canonical, parent_hash,
               gas_used, gas_limit, size, transaction_count, extra_data, base_fee_per_gas,
               transactions_root, state_root, receipts_root, logs_bloom, chain_id, retrieved_from
        FROM Blocks
        WHERE canonical = TRUE
        ORDER BY block_number DESC
        LIMIT $1 OFFSET $2
    `, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var block model.Block
		if err := rows.Scan(
			&block.BlockNumber, &block.BlockHash, &block.Timestamp, &block.MinerAddress,
			&block.Canonical, &block.ParentHash, &block.GasUsed, &block.GasLimit, &block.Size,
			&block.TransactionCount, &block.ExtraData, &block.BaseFeePerGas, &block.TransactionsRoot,
			&block.StateRoot, &block.ReceiptsRoot, &block.LogsBloom, &block.ChainID, &block.RetrievedFrom,
		); err != nil {
			return nil, err
		}
		blocks = append(blocks, block)
	}
	return blocks, nil
}

func (r *BlockRepoPG) GetBlockByNumber(blockNumber int64) (*model.Block, error) {
	var block model.Block
	err := r.db.QueryRow(context.Background(), `
        SELECT block_number, block_hash, timestamp, miner_address, canonical, parent_hash,
               gas_used, gas_limit, size, transaction_count, extra_data, base_fee_per_gas,
               transactions_root, state_root, receipts_root, logs_bloom, chain_id, retrieved_from
        FROM Blocks
        WHERE block_number = $1
    `, blockNumber).Scan(
		&block.BlockNumber, &block.BlockHash, &block.Timestamp, &block.MinerAddress,
		&block.Canonical, &block.ParentHash, &block.GasUsed, &block.GasLimit, &block.Size,
		&block.TransactionCount, &block.ExtraData, &block.BaseFeePerGas, &block.TransactionsRoot,
		&block.StateRoot, &block.ReceiptsRoot, &block.LogsBloom, &block.ChainID, &block.RetrievedFrom,
	)
	if err != nil {
		return nil, err
	}
	return &block, nil
}

func (r *BlockRepoPG) GetForkedBlocks(limit, offset int) ([]model.Block, error) {
	var blocks []model.Block
	rows, err := r.db.Query(context.Background(), `
        SELECT block_number, block_hash, timestamp, miner_address, canonical, parent_hash,
               gas_used, gas_limit, size, transaction_count, extra_data, base_fee_per_gas,
               transactions_root, state_root, receipts_root, logs_bloom, chain_id, retrieved_from
        FROM Blocks
        WHERE canonical = FALSE
        ORDER BY block_number DESC
        LIMIT $1 OFFSET $2
    `, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var block model.Block
		if err := rows.Scan(
			&block.BlockNumber, &block.BlockHash, &block.Timestamp, &block.MinerAddress,
			&block.Canonical, &block.ParentHash, &block.GasUsed, &block.GasLimit, &block.Size,
			&block.TransactionCount, &block.ExtraData, &block.BaseFeePerGas, &block.TransactionsRoot,
			&block.StateRoot, &block.ReceiptsRoot, &block.LogsBloom, &block.ChainID, &block.RetrievedFrom,
		); err != nil {
			return nil, err
		}
		blocks = append(blocks, block)
	}
	return blocks, nil
}

func (r *BlockRepoPG) GetBlockByHash(ctx context.Context, hash string) (*model.Block, error) {
	var block model.Block
	err := r.db.QueryRow(ctx, `
		SELECT block_number, block_hash, timestamp, miner_address, canonical, parent_hash,
		       gas_used, gas_limit, size, transaction_count, extra_data, base_fee_per_gas,
		       transactions_root, state_root, receipts_root, logs_bloom, chain_id, retrieved_from
		FROM blocks
		WHERE encode(block_hash, 'hex') = $1
	`, hash[2:]).Scan(
		&block.BlockNumber, &block.BlockHash, &block.Timestamp, &block.MinerAddress,
		&block.Canonical, &block.ParentHash, &block.GasUsed, &block.GasLimit, &block.Size,
		&block.TransactionCount, &block.ExtraData, &block.BaseFeePerGas, &block.TransactionsRoot,
		&block.StateRoot, &block.ReceiptsRoot, &block.LogsBloom, &block.ChainID, &block.RetrievedFrom,
	)
	if err != nil {
		return nil, err
	}
	return &block, nil
}
