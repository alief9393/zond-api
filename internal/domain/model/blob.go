package model

import "time"

type Blob struct {
	VersionedHash string    `db:"versioned_hash"`
	TxHash        string    `db:"tx_hash"`
	BlockNumber   int64     `db:"block_number"`
	Timestamp     time.Time `db:"timestamp"`
	BlobSender    string    `db:"blob_sender"`
	GasPrice      string    `db:"gas_price"`
	Size          int       `db:"size"`
	RetrievedFrom string    `db:"retrieved_from"`
}
