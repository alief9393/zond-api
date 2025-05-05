package service

import (
	"context"
	"zond-api/internal/api/dto"
	addressRepo "zond-api/internal/api/repository/address"
)

type AddressService interface {
	GetAddressBalance(ctx context.Context, address string) (*dto.AddressResponse, error)
	GetAddressTransactions(ctx context.Context, address string) (*dto.TransactionsResponse, error)
	GetAddressDetails(ctx context.Context, address string) (*dto.AddressDetailResponse, error) // ⬅️ tambahkan ini
}

type addressService struct {
	repo addressRepo.AddressRepo
}

func NewAddressService(repo addressRepo.AddressRepo) AddressService {
	return &addressService{repo: repo}
}

func (s *addressService) GetAddressBalance(ctx context.Context, address string) (*dto.AddressResponse, error) {
	return s.repo.GetAddressBalance(ctx, address)
}

func (s *addressService) GetAddressTransactions(ctx context.Context, address string) (*dto.TransactionsResponse, error) {
	return s.repo.GetAddressTransactions(ctx, address)
}

func (s *addressService) GetAddressDetails(ctx context.Context, address string) (*dto.AddressDetailResponse, error) {
	return s.repo.GetAddressDetails(ctx, address)
}
