package repository

import (
	"context"
	"zond-api/internal/api/dto"
	"zond-api/internal/domain/model"
)

type TransactionRepository interface {
	GetLatestTransactions(limit, offset int) ([]model.Transaction, error)
	GetTransactionByHash(txHash []byte) (*model.Transaction, error)
	GetTransactionsByBlockNumber(blockNumber int64) ([]model.Transaction, error)
	GetTransactionMetrics(ctx context.Context) (*dto.TransactionMetricsResponse, error)
	GetLatestTransactionsWithFilter(ctx context.Context, page, limit int, method, from, to string) ([]model.Transaction, error)
}
