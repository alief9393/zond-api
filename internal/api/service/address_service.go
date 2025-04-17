package service

import (
	"context"
	"zond-api/internal/api/dto"
	"zond-api/internal/api/repository"
)

type AddressService interface {
	GetAddressBalance(ctx context.Context, address string) (*dto.AddressResponse, error)
	GetAddressTransactions(ctx context.Context, address string) (*dto.TransactionsResponse, error)
}

type addressService struct {
	repo repository.AddressRepo
}

func NewAddressService(repo repository.AddressRepo) AddressService {
	return &addressService{repo: repo}
}

func (s *addressService) GetAddressBalance(ctx context.Context, address string) (*dto.AddressResponse, error) {
	return s.repo.GetAddressBalance(ctx, address)
}

func (s *addressService) GetAddressTransactions(ctx context.Context, address string) (*dto.TransactionsResponse, error) {
	return s.repo.GetAddressTransactions(ctx, address)
}
