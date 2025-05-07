package dto

import "time"

type BlockResponse struct {
	BlockNumber      int64     `json:"block_number"`
	BlockHash        string    `json:"block_hash"`
	Timestamp        time.Time `json:"timestamp"`
	MinerAddress     string    `json:"miner_address"`
	Canonical        bool      `json:"canonical"`
	ParentHash       string    `json:"parent_hash"`
	GasUsed          int64     `json:"gas_used"`
	GasLimit         int64     `json:"gas_limit"`
	Size             int64     `json:"size"`
	TransactionCount int       `json:"transaction_count"`
	ExtraData        string    `json:"extra_data"`
	BaseFeePerGas    int64     `json:"base_fee_per_gas"`
	TransactionsRoot string    `json:"transactions_root"`
	StateRoot        string    `json:"state_root"`
	ReceiptsRoot     string    `json:"receipts_root"`
	LogsBloom        string    `json:"logs_bloom"`
	ChainID          int64     `json:"chain_id"`
	RetrievedFrom    string    `json:"retrieved_from"`
	ReorgDepth       int       `json:"reorg_depth"`
}

type BlocksPaginatedResponse struct {
	Blocks     []BlockResponse `json:"blocks"`
	Pagination PaginationInfo  `json:"pagination"`
}
