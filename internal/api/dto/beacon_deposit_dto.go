package dto

import "time"

type BeaconDepositResponse struct {
	BlockNumber    int64     `json:"block_number" example:"123456"`
	TxHash         string    `json:"tx_hash" example:"0xabc123..."`
	Amount         string    `json:"amount" example:"32000000000"`
	ValidatorIndex int64     `json:"validator_index" example:"2940"`
	Timestamp      time.Time `json:"timestamp" example:"2024-05-06T14:20:00Z"`
	RetrievedFrom  string    `json:"retrieved_from" example:"geth"`
	LogIndex       int       `json:"log_index" example:"5"`
}

type BeaconDepositsPaginatedResponse struct {
	Deposits   []BeaconDepositResponse `json:"deposits"`
	Pagination PaginationInfo          `json:"pagination"`
}
