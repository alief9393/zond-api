package model

import "time"

type BundleTransaction struct {
	Hash        string    `db:"bundle_txn_hash"`
	BlockNumber int64     `db:"block_number"`
	Timestamp   time.Time `db:"timestamp"`
	Bundler     string    `db:"bundler"`
	EntryPoint  string    `db:"entry_point"`
	AATxnCount  int       `db:"aa_txns_count"`
	Amount      string    `db:"amount"`
	GasPrice    string    `db:"gas_price"`
}

type AATransaction struct {
	Hash        string    `db:"aa_txn_hash"`
	BundleHash  string    `db:"bundle_txn_hash"`
	Method      string    `db:"method"`
	BlockNumber int64     `db:"block_number"`
	Timestamp   time.Time `db:"timestamp"`
	From        string    `db:"from_address"`
	Bundler     string    `db:"bundler"`
	EntryPoint  string    `db:"entry_point"`
	GasPrice    string    `db:"gas_price"`
}
