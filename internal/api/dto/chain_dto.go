package dto

type ChainResponse struct {
	ChainID       int64  `json:"chain_id"`
	LatestBlock   int64  `json:"latest_block"`
	TotalBlocks   int64  `json:"total_blocks"`
	RetrievedFrom string `json:"retrieved_from"`
}
