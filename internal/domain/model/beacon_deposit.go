package model

import "time"

type BeaconDeposit struct {
	BlockNumber    int64     // SELECTED
	TxHash         []byte    // SELECTED
	Amount         string    // SELECTED
	ValidatorIndex int64     // SELECTED
	Timestamp      time.Time // SELECTED
	RetrievedFrom  string    // SELECTED
	LogIndex       int       // SELECTED
}
