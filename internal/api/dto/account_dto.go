package dto

type AccountResponse struct {
	Address    string  `json:"address"`
	NameTag    string  `json:"name_tag"`
	Balance    float64 `json:"balance"`
	Percentage float64 `json:"percentage"`
	TxCount    int     `json:"tx_count"`
}

type AccountsPaginatedResponse struct {
	Accounts   []AccountResponse `json:"accounts"`
	Pagination PaginationInfo    `json:"pagination"`
}
