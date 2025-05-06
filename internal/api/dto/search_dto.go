package dto

type SearchResult struct {
	Type  string      `json:"type"` // "block", "transaction", "address"
	Value interface{} `json:"value"`
}
type Suggestion struct {
	Type  string `json:"type" example:"transaction"`
	Value string `json:"value" example:"0xabc123..."`
}

type SearchSuggestionsResponse struct {
	Suggestions []Suggestion `json:"suggestions"`
}
