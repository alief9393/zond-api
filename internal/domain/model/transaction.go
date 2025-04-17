package model

type Transaction struct {
	TxHash               []byte  `json:"tx_hash" db:"tx_hash"`
	BlockNumber          int64   `json:"block_number" db:"block_number"`
	FromAddress          []byte  `json:"from_address" db:"from_address"`
	ToAddress            []byte  `json:"to_address" db:"to_address"`
	Value                string  `json:"value" db:"value"`
	Gas                  int64   `json:"gas" db:"gas"`
	GasPrice             string  `json:"gas_price" db:"gas_price"`
	Type                 int     `json:"type" db:"type"`
	ChainID              int64   `json:"chain_id" db:"chain_id"`
	AccessList           []byte  `json:"access_list" db:"access_list"`
	MaxFeePerGas         *string `json:"max_fee_per_gas" db:"max_fee_per_gas"`
	MaxPriorityFeePerGas *string `json:"max_priority_fee_per_gas" db:"max_priority_fee_per_gas"`
	TransactionIndex     int     `json:"transaction_index" db:"transaction_index"`
	CumulativeGasUsed    int64   `json:"cumulative_gas_used" db:"cumulative_gas_used"`
	IsSuccessful         bool    `json:"is_successful" db:"is_successful"`
	RetrievedFrom        string  `json:"retrieved_from" db:"retrieved_from"`
	IsCanonical          bool    `json:"is_canonical" db:"is_canonical"`
}
