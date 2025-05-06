package dto

import "time"

type ReorgResponse struct {
	BlockNumber   int64     `json:"block_number" example:"1234567"`
	OldBlockHash  string    `json:"old_block_hash" example:"0xoldhash..."`
	NewBlockHash  string    `json:"new_block_hash" example:"0xnewhash..."`
	Depth         int       `json:"depth" example:"2"`
	Timestamp     time.Time `json:"timestamp" example:"2024-05-06T12:00:00Z"`
	ChainID       int64     `json:"chain_id" example:"1"`
	RetrievedFrom string    `json:"retrieved_from" example:"nethermind"`
}

type ReorgsResponse struct {
	Reorgs []ReorgResponse `json:"reorgs"`
}
