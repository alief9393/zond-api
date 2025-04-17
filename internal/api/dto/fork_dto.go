package dto

import "time"

type ForkResponse struct {
	ForkName      string    `json:"fork_name"`
	BlockNumber   int64     `json:"block_number"`
	Timestamp     time.Time `json:"timestamp"`
	ChainID       int64     `json:"chain_id"`
	RetrievedFrom string    `json:"retrieved_from"`
}

type ForksResponse struct {
	Forks []ForkResponse `json:"forks"`
}
