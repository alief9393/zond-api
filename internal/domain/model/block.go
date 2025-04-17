package model

import "time"

type Block struct {
	BlockNumber      int64     `json:"block_number" db:"block_number"`
	BlockHash        []byte    `json:"block_hash" db:"block_hash"`
	Timestamp        time.Time `json:"timestamp" db:"timestamp"`
	MinerAddress     []byte    `json:"miner_address" db:"miner_address"`
	Canonical        bool      `json:"canonical" db:"canonical"`
	ParentHash       []byte    `json:"parent_hash" db:"parent_hash"`
	GasUsed          string    `json:"gas_used" db:"gas_used"`
	GasLimit         string    `json:"gas_limit" db:"gas_limit"`
	Size             int       `json:"size" db:"size"`
	TransactionCount int       `json:"transaction_count" db:"transaction_count"`
	ExtraData        []byte    `json:"extra_data" db:"extra_data"`
	BaseFeePerGas    *int64    `json:"base_fee_per_gas" db:"base_fee_per_gas"`
	TransactionsRoot []byte    `json:"transactions_root" db:"transactions_root"`
	StateRoot        []byte    `json:"state_root" db:"state_root"`
	ReceiptsRoot     []byte    `json:"receipts_root" db:"receipts_root"`
	LogsBloom        []byte    `json:"logs_bloom" db:"logs_bloom"`
	ChainID          int64     `json:"chain_id" db:"chain_id"`
	RetrievedFrom    string    `json:"retrieved_from" db:"retrieved_from"`
}
