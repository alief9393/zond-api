package dto

type TransactionResponse struct {
	TxHash               string  `json:"tx_hash" example:"0xabc123..."`
	BlockNumber          int64   `json:"block_number" example:"123456"`
	FromAddress          string  `json:"from_address" example:"0x1111..."`
	ToAddress            string  `json:"to_address" example:"0x2222..."`
	Value                string  `json:"value" example:"1000000000000000000"`
	Gas                  int64   `json:"gas" example:"21000"`
	GasPrice             string  `json:"gas_price" example:"5000000000"`
	Type                 int     `json:"type" example:"2"`
	ChainID              int64   `json:"chain_id" example:"1"`
	AccessList           string  `json:"access_list" example:"[]"`
	MaxFeePerGas         *string `json:"max_fee_per_gas" example:"6000000000"`
	MaxPriorityFeePerGas *string `json:"max_priority_fee_per_gas" example:"2000000000"`
	TransactionIndex     int     `json:"transaction_index" example:"0"`
	CumulativeGasUsed    int64   `json:"cumulative_gas_used" example:"21000"`
	IsSuccessful         bool    `json:"is_successful" example:"true"`
	RetrievedFrom        string  `json:"retrieved_from" example:"geth"`
	IsCanonical          bool    `json:"is_canonical" example:"true"`
}

type TransactionsResponse struct {
	Transactions []TransactionResponse `json:"transactions"`
}

type TransactionsPaginatedResponse struct {
	Transactions []TransactionResponse `json:"transactions"`
	Pagination   PaginationInfo        `json:"pagination"`
}

type TransactionMetricsResponse struct {
	Transactions24h       int     `json:"transactions_24h" example:"34256"`
	PendingTransactions1h int     `json:"pending_transactions_1h" example:"135"`
	NetworkFeeETH24h      float64 `json:"network_fee_eth_24h" example:"25.67"`
	AvgFeeUSD24h          float64 `json:"avg_fee_usd_24h" example:"1.25"`
}

type DailyTransactionStat struct {
	Date  string `json:"date" example:"2025-05-07"`
	Count int    `json:"count" example:"12345"`
}

type DailyTransactionStatsResponse struct {
	Data []DailyTransactionStat `json:"data"`
}

type TPSStatResponse struct {
	AverageTPS float64 `json:"average_tps" example:"12.45"`
}

type DailyFeeStat struct {
	Date          string  `json:"date" example:"2025-05-07"`
	NetworkFeeETH float64 `json:"network_fee_eth" example:"1.23"`
	AvgFeeUSD     float64 `json:"avg_fee_usd" example:"0.56"`
}

type DailyFeeStatsResponse struct {
	Data []DailyFeeStat `json:"data"`
}
