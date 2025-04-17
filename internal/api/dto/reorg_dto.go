package dto

import "time"

type ReorgResponse struct {
	BlockNumber   int64     `json:"block_number"`
	OldBlockHash  string    `json:"old_block_hash"`
	NewBlockHash  string    `json:"new_block_hash"`
	Depth         int       `json:"depth"`
	Timestamp     time.Time `json:"timestamp"`
	ChainID       int64     `json:"chain_id"`
	RetrievedFrom string    `json:"retrieved_from"`
}

type ReorgsResponse struct {
	Reorgs []ReorgResponse `json:"reorgs"`
}
