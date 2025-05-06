package beacon_withdrawal

import (
	"context"
	"zond-api/internal/domain/model"
)

type BeaconWithdrawalRepository interface {
	GetBeaconWithdrawals(ctx context.Context, limit, offset int) ([]model.BeaconWithdrawal, error)
	CountBeaconWithdrawals(ctx context.Context) (int, error)
}
