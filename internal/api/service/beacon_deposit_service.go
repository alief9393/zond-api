package service

import (
	"context"
	"fmt"
	"zond-api/internal/api/dto"
	"zond-api/internal/api/repository/beacon_deposit"
)

type BeaconDepositService interface {
	GetBeaconDeposits(ctx context.Context, page, limit int) ([]dto.BeaconDepositResponse, int, error)
}

type beaconDepositService struct {
	repo beacon_deposit.BeaconDepositRepository
}

func NewBeaconDepositService(repo beacon_deposit.BeaconDepositRepository) BeaconDepositService {
	return &beaconDepositService{repo: repo}
}

func (s *beaconDepositService) GetBeaconDeposits(ctx context.Context, page, limit int) ([]dto.BeaconDepositResponse, int, error) {
	offset := (page - 1) * limit

	models, err := s.repo.GetBeaconDeposits(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.CountBeaconDeposits(ctx)
	if err != nil {
		return nil, 0, err
	}

	var responses []dto.BeaconDepositResponse
	for _, m := range models {
		responses = append(responses, dto.BeaconDepositResponse{
			BlockNumber:    m.BlockNumber,
			TxHash:         fmt.Sprintf("0x%x", m.TxHash),
			Amount:         m.Amount,
			ValidatorIndex: m.ValidatorIndex,
			Timestamp:      m.Timestamp,
			RetrievedFrom:  m.RetrievedFrom,
			LogIndex:       m.LogIndex,
		})
	}

	return responses, total, nil
}
