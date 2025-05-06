package service

import (
	"context"
	"fmt"

	"zond-api/internal/api/dto"
	"zond-api/internal/api/repository/beacon_withdrawal"
)

type BeaconWithdrawalService interface {
	GetBeaconWithdrawals(ctx context.Context, page, limit int) ([]dto.BeaconWithdrawalResponse, int, error)
}

type beaconWithdrawalService struct {
	repo beacon_withdrawal.BeaconWithdrawalRepository
}

func NewBeaconWithdrawalService(repo beacon_withdrawal.BeaconWithdrawalRepository) BeaconWithdrawalService {
	return &beaconWithdrawalService{repo: repo}
}

func (s *beaconWithdrawalService) GetBeaconWithdrawals(ctx context.Context, page, limit int) ([]dto.BeaconWithdrawalResponse, int, error) {
	offset := (page - 1) * limit
	models, err := s.repo.GetBeaconWithdrawals(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.CountBeaconWithdrawals(ctx)
	if err != nil {
		return nil, 0, err
	}

	var responses []dto.BeaconWithdrawalResponse
	for _, m := range models {
		responses = append(responses, dto.BeaconWithdrawalResponse{
			BlockNumber:    m.BlockNumber,
			TxHash:         fmt.Sprintf("0x%x", m.TxHash),
			Amount:         m.Amount,
			ValidatorIndex: m.ValidatorIndex,
			Timestamp:      m.Timestamp,
			RetrievedFrom:  m.RetrievedFrom,
			LogIndex:       int64(m.LogIndex),
		})
	}

	return responses, total, nil
}
