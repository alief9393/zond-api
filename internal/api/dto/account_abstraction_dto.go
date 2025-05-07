package dto

import "time"

type BundleTransactionResponse struct {
	Hash        string    `json:"bundle_txn_hash"`
	BlockNumber int64     `json:"block_number"`
	Timestamp   time.Time `json:"timestamp"`
	Bundler     string    `json:"bundler"`
	EntryPoint  string    `json:"entry_point"`
	AATxnCount  int       `json:"aa_txns_count"`
	Amount      string    `json:"amount"`
	GasPrice    string    `json:"gas_price"`
}

type AATransactionResponse struct {
	Hash        string    `json:"aa_txn_hash"`
	BundleHash  string    `json:"bundle_txn_hash"`
	Method      string    `json:"method"`
	BlockNumber int64     `json:"block_number"`
	Timestamp   time.Time `json:"timestamp"`
	From        string    `json:"from"`
	Bundler     string    `json:"bundler"`
	EntryPoint  string    `json:"entry_point"`
	GasPrice    string    `json:"gas_price"`
}

type AccountAbstractionResponse struct {
	Bundles []BundleTransactionResponse `json:"bundles"`
	AATxns  []AATransactionResponse     `json:"aa_transactions"`
}

type AccountAbstractionPaginatedResponse struct {
	Data       AccountAbstractionResponse `json:"data"`
	Pagination PaginationInfo             `json:"pagination"`
}

type BundlesPaginatedResponse struct {
	Bundles    []BundleTransactionResponse `json:"bundles"`
	Pagination PaginationInfo              `json:"pagination"`
}

type AATransactionsPaginatedResponse struct {
	AATransactions []AATransactionResponse `json:"aa_transactions"`
	Pagination     PaginationInfo          `json:"pagination"`
}
