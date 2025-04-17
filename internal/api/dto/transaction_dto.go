package dto

type TransactionResponse struct {
	TxHash               string  `json:"tx_hash"`
	BlockNumber          int64   `json:"block_number"`
	FromAddress          string  `json:"from_address"`
	ToAddress            string  `json:"to_address"`
	Value                string  `json:"value"`
	Gas                  int64   `json:"gas"`
	GasPrice             string  `json:"gas_price"`
	Type                 int     `json:"type"`
	ChainID              int64   `json:"chain_id"`
	AccessList           string  `json:"access_list"`
	MaxFeePerGas         *string `json:"max_fee_per_gas"`
	MaxPriorityFeePerGas *string `json:"max_priority_fee_per_gas"`
	TransactionIndex     int     `json:"transaction_index"`
	CumulativeGasUsed    int64   `json:"cumulative_gas_used"`
	IsSuccessful         bool    `json:"is_successful"`
	RetrievedFrom        string  `json:"retrieved_from"`
	IsCanonical          bool    `json:"is_canonical"`
}

type TransactionsResponse struct {
	Transactions []TransactionResponse `json:"transactions"`
}
