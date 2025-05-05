package dto

type AddressDTO struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
}

type TransactionDTO struct {
	Hash        string `json:"hash"`
	BlockNumber uint64 `json:"block_number"`
	From        string `json:"from"`
	To          string `json:"to"`
	Value       string `json:"value"`
}

type AddressResponse struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
}

type TopAddressesResponse struct {
	Addresses []AddressResponse `json:"addresses"`
}

type AddressDetailResponse struct {
	Address          string  `json:"address"`
	Balance          float64 `json:"balance"`
	TransactionCount int     `json:"transaction_count"`
	TotalGasUsed     int64   `json:"total_gas_used"`
}
