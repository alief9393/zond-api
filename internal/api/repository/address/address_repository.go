package repository

import (
	"context"
	"zond-api/internal/api/dto"
)

type AddressRepository interface {
	GetAddressBalance(ctx context.Context, address string) (*dto.AddressResponse, error)
	GetAddressTransactions(ctx context.Context, address string) (*dto.TransactionsResponse, error)
	GetTopAddresses(ctx context.Context, limit int) (*dto.TopAddressesResponse, error)
	GetAddressDetails(ctx context.Context, address string) (*dto.AddressDetailResponse, error)
}
