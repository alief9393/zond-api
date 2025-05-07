package service

import (
	"context"
	"zond-api/internal/api/dto"
	"zond-api/internal/api/repository/account_abstraction"
)

type AccountAbstractionService interface {
	GetAccountAbstraction(ctx context.Context, page, limit int) (dto.AccountAbstractionResponse, int, error)
	GetBundlesOnly(ctx context.Context, page, limit int) ([]dto.BundleTransactionResponse, int, error)
	GetAATransactionsOnly(ctx context.Context, page, limit int) ([]dto.AATransactionResponse, int, error)
}

type accountAbstractionService struct {
	repo account_abstraction.AccountAbstractionRepository
}

func NewAccountAbstractionService(repo account_abstraction.AccountAbstractionRepository) AccountAbstractionService {
	return &accountAbstractionService{repo: repo}
}

func (s *accountAbstractionService) GetAccountAbstraction(ctx context.Context, page, limit int) (dto.AccountAbstractionResponse, int, error) {
	offset := (page - 1) * limit

	bundles, total, err := s.repo.GetBundleTransactions(limit, offset)
	if err != nil {
		return dto.AccountAbstractionResponse{}, 0, err
	}

	aaTxns, _, err := s.repo.GetAATransactions(limit, offset)
	if err != nil {
		return dto.AccountAbstractionResponse{}, 0, err
	}

	var res dto.AccountAbstractionResponse
	for _, b := range bundles {
		res.Bundles = append(res.Bundles, dto.BundleTransactionResponse{
			Hash:        b.Hash,
			BlockNumber: b.BlockNumber,
			Timestamp:   b.Timestamp,
			Bundler:     b.Bundler,
			EntryPoint:  b.EntryPoint,
			AATxnCount:  b.AATxnCount,
			Amount:      b.Amount,
			GasPrice:    b.GasPrice,
		})
	}

	for _, a := range aaTxns {
		res.AATxns = append(res.AATxns, dto.AATransactionResponse{
			Hash:        a.Hash,
			BundleHash:  a.BundleHash,
			Method:      a.Method,
			BlockNumber: a.BlockNumber,
			Timestamp:   a.Timestamp,
			From:        a.From,
			Bundler:     a.Bundler,
			EntryPoint:  a.EntryPoint,
			GasPrice:    a.GasPrice,
		})
	}

	return res, total, nil
}

func (s *accountAbstractionService) GetBundlesOnly(ctx context.Context, page, limit int) ([]dto.BundleTransactionResponse, int, error) {
	offset := (page - 1) * limit
	bundles, total, err := s.repo.GetBundleTransactions(limit, offset)
	if err != nil {
		return nil, 0, err
	}

	var res []dto.BundleTransactionResponse
	for _, b := range bundles {
		res = append(res, dto.BundleTransactionResponse{
			Hash:        b.Hash,
			BlockNumber: b.BlockNumber,
			Timestamp:   b.Timestamp,
			Bundler:     b.Bundler,
			EntryPoint:  b.EntryPoint,
			AATxnCount:  b.AATxnCount,
			Amount:      b.Amount,
			GasPrice:    b.GasPrice,
		})
	}
	return res, total, nil
}

func (s *accountAbstractionService) GetAATransactionsOnly(ctx context.Context, page, limit int) ([]dto.AATransactionResponse, int, error) {
	offset := (page - 1) * limit
	txns, total, err := s.repo.GetAATransactions(limit, offset)
	if err != nil {
		return nil, 0, err
	}

	var res []dto.AATransactionResponse
	for _, t := range txns {
		res = append(res, dto.AATransactionResponse{
			Hash:        t.Hash,
			BundleHash:  t.BundleHash,
			Method:      t.Method,
			BlockNumber: t.BlockNumber,
			Timestamp:   t.Timestamp,
			From:        t.From,
			Bundler:     t.Bundler,
			EntryPoint:  t.EntryPoint,
			GasPrice:    t.GasPrice,
		})
	}
	return res, total, nil
}
