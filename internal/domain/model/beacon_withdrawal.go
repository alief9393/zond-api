package model

import "time"

type BeaconWithdrawal struct {
	BlockNumber    int64     `db:"block_number"`
	TxHash         []byte    `db:"tx_hash"`
	ValidatorIndex int64     `db:"validator_index"`
	Amount         string    `db:"amount"`
	Timestamp      time.Time `db:"timestamp"`
	RetrievedFrom  string    `db:"retrieved_from"`
	LogIndex       int       `db:"log_index"`
}
