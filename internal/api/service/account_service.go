package service

import (
	"context"
	"fmt"
	"zond-api/internal/api/dto"
	accountRepo "zond-api/internal/api/repository/account"
)

type AccountService struct {
	repo accountRepo.AccountRepository
}

func NewAccountService(repo accountRepo.AccountRepository) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) GetTopAccounts(ctx context.Context, page, limit int) (dto.AccountsPaginatedResponse, error) {
	accounts, total, err := s.repo.GetTopAccounts(ctx, page, limit)
	if err != nil {
		return dto.AccountsPaginatedResponse{}, err
	}

	var responses []dto.AccountResponse
	for _, acc := range accounts {
		responses = append(responses, dto.AccountResponse{
			Address:    fmt.Sprintf("0x%x", acc.Address),
			NameTag:    acc.NameTag,
			Balance:    acc.Balance,
			Percentage: acc.Percentage,
			TxCount:    acc.TxCount,
		})
	}

	return dto.AccountsPaginatedResponse{
		Accounts: responses,
		Pagination: dto.PaginationInfo{
			Total: total,
			Page:  page,
			Limit: limit,
		},
	}, nil
}
