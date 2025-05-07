package dto

import "time"

type BlobResponse struct {
	VersionedHash string    `json:"versioned_hash"`
	TxHash        string    `json:"tx_hash"`
	BlockNumber   int64     `json:"block_number"`
	Timestamp     time.Time `json:"timestamp"`
	BlobSender    string    `json:"blob_sender"`
	GasPrice      string    `json:"gas_price"`
	Size          int       `json:"size"` // in bytes
	RetrievedFrom string    `json:"retrieved_from"`
}

type BlobsPaginatedResponse struct {
	Blobs      []BlobResponse `json:"blobs"`
	Pagination PaginationInfo `json:"pagination"`
}
