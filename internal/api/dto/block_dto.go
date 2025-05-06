package dto

import "time"

type BlockResponse struct {
	BlockNumber      int64     `json:"block_number" example:"123456"`
	BlockHash        string    `json:"block_hash" example:"0xabc123..."`
	Timestamp        time.Time `json:"timestamp" example:"2024-05-06T10:00:00Z"`
	MinerAddress     string    `json:"miner_address" example:"0x1a2b..."`
	Canonical        bool      `json:"canonical" example:"true"`
	ParentHash       string    `json:"parent_hash" example:"0xdeadbeef..."`
	GasUsed          string    `json:"gas_used" example:"21000"`
	GasLimit         string    `json:"gas_limit" example:"30000000"`
	Size             int       `json:"size" example:"1200"`
	TransactionCount int       `json:"transaction_count" example:"15"`
	ExtraData        string    `json:"extra_data" example:"0x..."`
	BaseFeePerGas    *int64    `json:"base_fee_per_gas" example:"1000000000"`
	TransactionsRoot string    `json:"transactions_root" example:"0xroot..."`
	StateRoot        string    `json:"state_root" example:"0xstate..."`
	ReceiptsRoot     string    `json:"receipts_root" example:"0xreceipt..."`
	LogsBloom        string    `json:"logs_bloom" example:"0x..."`
	ChainID          int64     `json:"chain_id" example:"1"`
	RetrievedFrom    string    `json:"retrieved_from" example:"geth"`
}

type BlocksResponse struct {
	Blocks []BlockResponse `json:"blocks"`
}
