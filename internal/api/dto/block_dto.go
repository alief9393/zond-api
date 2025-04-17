package dto

import "time"

type BlockResponse struct {
	BlockNumber      int64     `json:"block_number"`
	BlockHash        string    `json:"block_hash"`
	Timestamp        time.Time `json:"timestamp"`
	MinerAddress     string    `json:"miner_address"`
	Canonical        bool      `json:"canonical"`
	ParentHash       string    `json:"parent_hash"`
	GasUsed          string    `json:"gas_used"`
	GasLimit         string    `json:"gas_limit"`
	Size             int       `json:"size"`
	TransactionCount int       `json:"transaction_count"`
	ExtraData        string    `json:"extra_data"`
	BaseFeePerGas    *int64    `json:"base_fee_per_gas"`
	TransactionsRoot string    `json:"transactions_root"`
	StateRoot        string    `json:"state_root"`
	ReceiptsRoot     string    `json:"receipts_root"`
	LogsBloom        string    `json:"logs_bloom"`
	ChainID          int64     `json:"chain_id"`
	RetrievedFrom    string    `json:"retrieved_from"`
}

type BlocksResponse struct {
	Blocks []BlockResponse `json:"blocks"`
}
