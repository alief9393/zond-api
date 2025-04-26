package dto

type SearchResult struct {
	Type  string      `json:"type"` // "block", "transaction", "address"
	Value interface{} `json:"value"`
}
