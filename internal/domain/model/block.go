package model

import "time"

type Block struct {
	BlockNumber      int64     `json:"block_number"`
	BlockHash        []byte    `json:"block_hash"`
	Timestamp        time.Time `json:"timestamp"`
	MinerAddress     []byte    `json:"miner_address"`
	Canonical        bool      `json:"canonical"`
	ParentHash       []byte    `json:"parent_hash"`
	GasUsed          int64     `json:"gas_used"`
	GasLimit         int64     `json:"gas_limit"`
	Size             int64     `json:"size"`
	TransactionCount int       `json:"transaction_count"`
	ExtraData        []byte    `json:"extra_data"`
	BaseFeePerGas    int64     `json:"base_fee_per_gas"`
	TransactionsRoot []byte    `json:"transactions_root"`
	StateRoot        []byte    `json:"state_root"`
	ReceiptsRoot     []byte    `json:"receipts_root"`
	LogsBloom        []byte    `json:"logs_bloom"`
	ChainID          int64     `json:"chain_id"`
	RetrievedFrom    string    `json:"retrieved_from"`
	Slot             int64     `json:"slot"`
	RewardETH        float64   `json:"reward_eth"`
	BurntFeesETH     float64   `json:"burnt_fees_eth"`
	ReorgDepth       int       `json:"reorg_depth"`
}
