package service

import (
	"context"
	"encoding/hex"
	"fmt"

	"zond-api/internal/api/dto"
	txRepo "zond-api/internal/api/repository/transaction"
	"zond-api/internal/domain/model"
)

type TransactionService struct {
	repo txRepo.TransactionRepository
}

func NewTransactionService(repo txRepo.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) GetLatestTransactions(limit, offset int) (dto.TransactionsResponse, error) {
	transactions, err := s.repo.GetLatestTransactions(limit, offset)
	if err != nil {
		return dto.TransactionsResponse{}, err
	}

	var txResponses []dto.TransactionResponse
	for _, tx := range transactions {
		toAddress := ""
		if len(tx.ToAddress) > 0 {
			toAddress = fmt.Sprintf("0x%x", tx.ToAddress)
		}

		txResponses = append(txResponses, dto.TransactionResponse{
			TxHash:               fmt.Sprintf("0x%x", tx.TxHash),
			BlockNumber:          tx.BlockNumber,
			FromAddress:          fmt.Sprintf("0x%x", tx.FromAddress),
			ToAddress:            toAddress,
			Value:                tx.Value,
			Gas:                  tx.Gas,
			GasPrice:             tx.GasPrice,
			Type:                 tx.Type,
			ChainID:              tx.ChainID,
			AccessList:           string(tx.AccessList),
			MaxFeePerGas:         tx.MaxFeePerGas,
			MaxPriorityFeePerGas: tx.MaxPriorityFeePerGas,
			TransactionIndex:     tx.TransactionIndex,
			CumulativeGasUsed:    tx.CumulativeGasUsed,
			IsSuccessful:         tx.IsSuccessful,
			RetrievedFrom:        tx.RetrievedFrom,
			IsCanonical:          tx.IsCanonical,
		})
	}
	return dto.TransactionsResponse{Transactions: txResponses}, nil
}

func (s *TransactionService) GetTransactionByHash(txHash string) (*dto.TransactionResponse, error) {
	if len(txHash) > 2 && txHash[:2] == "0x" {
		txHash = txHash[2:]
	}
	hashBytes, err := hex.DecodeString(txHash)
	if err != nil {
		return nil, fmt.Errorf("invalid transaction hash: %w", err)
	}

	tx, err := s.repo.GetTransactionByHash(hashBytes)
	if err != nil {
		return nil, err
	}

	toAddress := ""
	if len(tx.ToAddress) > 0 {
		toAddress = fmt.Sprintf("0x%x", tx.ToAddress)
	}

	return &dto.TransactionResponse{
		TxHash:               fmt.Sprintf("0x%x", tx.TxHash),
		BlockNumber:          tx.BlockNumber,
		FromAddress:          fmt.Sprintf("0x%x", tx.FromAddress),
		ToAddress:            toAddress,
		Value:                tx.Value,
		Gas:                  tx.Gas,
		GasPrice:             tx.GasPrice,
		Type:                 tx.Type,
		ChainID:              tx.ChainID,
		AccessList:           string(tx.AccessList),
		MaxFeePerGas:         tx.MaxFeePerGas,
		MaxPriorityFeePerGas: tx.MaxPriorityFeePerGas,
		TransactionIndex:     tx.TransactionIndex,
		CumulativeGasUsed:    tx.CumulativeGasUsed,
		IsSuccessful:         tx.IsSuccessful,
		RetrievedFrom:        tx.RetrievedFrom,
		IsCanonical:          tx.IsCanonical,
	}, nil
}

func (s *TransactionService) GetTransactionsByBlockNumber(blockNumber int64) ([]dto.TransactionResponse, error) {
	txs, err := s.repo.GetTransactionsByBlockNumber(blockNumber)
	if err != nil {
		return nil, err
	}

	var responses []dto.TransactionResponse
	for _, tx := range txs {
		toAddr := ""
		if len(tx.ToAddress) > 0 {
			toAddr = fmt.Sprintf("0x%x", tx.ToAddress)
		}
		responses = append(responses, dto.TransactionResponse{
			TxHash:               fmt.Sprintf("0x%x", tx.TxHash),
			BlockNumber:          tx.BlockNumber,
			FromAddress:          fmt.Sprintf("0x%x", tx.FromAddress),
			ToAddress:            toAddr,
			Value:                tx.Value,
			Gas:                  tx.Gas,
			GasPrice:             tx.GasPrice,
			Type:                 tx.Type,
			ChainID:              tx.ChainID,
			AccessList:           string(tx.AccessList),
			MaxFeePerGas:         tx.MaxFeePerGas,
			MaxPriorityFeePerGas: tx.MaxPriorityFeePerGas,
			TransactionIndex:     tx.TransactionIndex,
			CumulativeGasUsed:    tx.CumulativeGasUsed,
			IsSuccessful:         tx.IsSuccessful,
			RetrievedFrom:        tx.RetrievedFrom,
			IsCanonical:          tx.IsCanonical,
		})
	}
	return responses, nil
}

func (s *TransactionService) GetTransactionMetrics(ctx context.Context) (*dto.TransactionMetricsResponse, error) {
	return s.repo.GetTransactionMetrics(ctx)
}

func (s *TransactionService) GetLatestTransactionsWithFilter(ctx context.Context, page, limit int, method, from, to string) ([]model.Transaction, error) {
	return s.repo.GetLatestTransactionsWithFilter(ctx, page, limit, method, from, to)
}
