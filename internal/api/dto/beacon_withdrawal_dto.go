package dto

import "time"

type BeaconWithdrawalResponse struct {
	BlockNumber    int64     `json:"block_number" example:"1200456"`
	TxHash         string    `json:"tx_hash" example:"0xabc123..."`
	ValidatorIndex int64     `json:"validator_index" example:"389"`
	Amount         string    `json:"amount" example:"34000000000"`
	Timestamp      time.Time `json:"timestamp" example:"2024-05-06T14:20:00Z"`
	RetrievedFrom  string    `json:"retrieved_from" example:"geth"`
	LogIndex       int64     `json:"log_index" example:"7"`
}

type BeaconWithdrawalsPaginatedResponse struct {
	Withdrawals []BeaconWithdrawalResponse `json:"withdrawals"`
	Pagination  PaginationInfo             `json:"pagination"`
}
