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
	CountTransactionsWithFilter(ctx context.Context, method, from, to string) (int, error)
	GetPendingTransactions(ctx context.Context, page, limit int, method, from, to string) ([]model.Transaction, error)
	CountPendingTransactions(ctx context.Context, method, from, to string) (int, error)
	GetContractTransactions(ctx context.Context, page, limit int, method, from, to string) ([]model.Transaction, error)
	CountContractTransactions(ctx context.Context, method, from, to string) (int, error)
}
