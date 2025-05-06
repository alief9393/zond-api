package beacon_deposit

import (
	"context"
	"zond-api/internal/domain/model"
)

type BeaconDepositRepository interface {
	GetBeaconDeposits(ctx context.Context, page, limit int) ([]model.BeaconDeposit, error)
	CountBeaconDeposits(ctx context.Context) (int, error)
}
