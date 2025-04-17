package repository

import "zond-api/internal/domain/model"

type TransactionRepository interface {
	GetLatestTransactions(limit, offset int) ([]model.Transaction, error)
	GetTransactionByHash(txHash []byte) (*model.Transaction, error)
}
