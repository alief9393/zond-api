package dto

type SearchResult struct {
	Type  string      `json:"type"` // "block", "transaction", "address"
	Value interface{} `json:"value"`
}
type Suggestion struct {
	Type  string `json:"type"`  // e.g., "address", "transaction", "block", etc.
	Value string `json:"value"` // e.g., hash, address, number
}

type SearchSuggestionsResponse struct {
	Suggestions []Suggestion `json:"suggestions"`
}
