package dto

import "time"

type ForkResponse struct {
	ForkName      string    `json:"fork_name" example:"Shanghai"`
	BlockNumber   int64     `json:"block_number" example:"17890000"`
	Timestamp     time.Time `json:"timestamp" example:"2024-05-06T10:00:00Z"`
	ChainID       int64     `json:"chain_id" example:"1"`
	RetrievedFrom string    `json:"retrieved_from" example:"geth"`
}

type ForksResponse struct {
	Forks []ForkResponse `json:"forks"`
}
