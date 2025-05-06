package dto

type ChainResponse struct {
	ChainID       int64  `json:"chain_id" example:"1"`
	LatestBlock   int64  `json:"latest_block" example:"1245678"`
	TotalBlocks   int64  `json:"total_blocks" example:"1245678"`
	RetrievedFrom string `json:"retrieved_from" example:"geth"`
}
